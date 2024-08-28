package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var scrappingBaseUrl string = "https://www.saramin.co.kr/zf_user/jobs/list/job-category?"
var jobDetailUrl string = "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx="

type ExtractJob struct {
	id        string
	company   string
	title     string
	workPlace string
	career    string
	summary   []string
}

// Scrape web page
func Scrape(searchKeyword string) {
	start := time.Now()
	startPage := 1

	extractJobChannel := make(chan []ExtractJob)

	totalPages := getTotalPageCount(startPage, searchKeyword)

	var jobs []ExtractJob
	for i := 1; i <= totalPages; i++ {
		go getPageByPageNum(i, searchKeyword, extractJobChannel)
	}
	for i := 1; i <= totalPages; i++ {
		jobs = append(jobs, <-extractJobChannel...)
	}

	writeJob(searchKeyword, jobs)
	end := time.Since(start)
	fmt.Println("ended....:", end)
}

// CleanString string sanitize
func CleanString(txt string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(txt)), " ")
}

func writeJob(searchKeyword string, jobs []ExtractJob) {

	file, err := os.Create(searchKeyword + "_jobs.csv")
	checkErr(err)
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)

	headers := []string{"Link", "Company", "Title", "WorkPlace", "Career", "summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	done := make(chan []string)

	for _, job := range jobs {
		go func(job ExtractJob) {
			done <- []string{jobDetailUrl + job.id, job.company, job.title, job.workPlace, job.career, strings.Join(job.summary, " ")}
		}(job)
	}

	for i := 0; i < len(jobs); i++ {

		w.Write(<-done)
	}

	// for _, job := range jobs {
	// 	jobSlice := []string{jobDetailUrl + job.id, job.company, job.title, job.workPlace, job.career, strings.Join(job.summary, " ")}
	// 	wErr := w.Write(jobSlice)
	// 	checkErr(wErr)
	// }

	fmt.Println("Done. count", len(jobs))
	defer w.Flush()
	defer file.Close()
}

func getPageByPageNum(pageNum int, searchKeyword string, extractJobChannel chan<- []ExtractJob) {
	c := make(chan ExtractJob)
	pageUrl := scrappingBaseUrl + "page=" + strconv.Itoa(pageNum) + "&cat_kewd=84%2C87%2C2232&searchType=search&searchword=" + searchKeyword + "&search_optional_item=y&search_done=y&panel_count=y&preview=y&isAjaxRequest=0&page_count=50&sort=RL&type=job-category&is_param=1&isSearchResultEmpty=1&isSectionHome=0&searchParamCount=2#searchTitle"
	fmt.Println("Request page :", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	var extractJobs []ExtractJob
	rawJobs := doc.Find(".list_recruiting > div.list_body > .list_item")
	rawJobs.Each(func(i int, card *goquery.Selection) {
		go makeExtractJob(card, c)

	})

	for i := 0; i < rawJobs.Length(); i++ {
		job := <-c
		extractJobs = append(extractJobs, job)
	}
	extractJobChannel <- extractJobs
	defer res.Body.Close() // 어디에 배치하던 결국 마지막에 실행

}

func makeExtractJob(card *goquery.Selection, c chan<- ExtractJob) {
	recIdx, _ := card.Attr("id")
	title := card.Find(".job_tit>a>span").Text()

	companyNm := card.Find(".box_item > .company_nm > .str_tit").Text()
	recruitInfo := card.Find(".box_item > .recruit_info > ul")

	workPlace := recruitInfo.Find(".work_place").Text()
	career := recruitInfo.Find(".career").Text()

	var sectors []string
	card.Find(".box_item > .notification_info > .job_meta > .job_sector>span").Each(func(i int, sector *goquery.Selection) {
		trimSector := CleanString(sector.Text())
		sectors = append(sectors, trimSector)

	})

	extractJob := ExtractJob{
		id:        strings.Split(recIdx, "-")[1],
		company:   CleanString(companyNm),
		title:     CleanString(title),
		workPlace: workPlace,
		career:    career,
		summary:   sectors,
	}
	c <- extractJob
}

func getTotalPageCount(startPage int, searchKeyword string) int {

	res, err := http.Get(scrappingBaseUrl + "page=" + strconv.Itoa(startPage) + "&cat_kewd=84%2C87%2C2232&searchType=search&searchword=" + searchKeyword + "&search_optional_item=y&search_done=y&panel_count=y&preview=y&isAjaxRequest=0&page_count=50&sort=RL&type=job-category&is_param=1&isSearchResultEmpty=1&isSectionHome=0&searchParamCount=2#searchTitle")
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	isNextButton := false
	buttons := doc.Find(".PageBox > button")
	buttons.Each(func(i int, button *goquery.Selection) {

		buttonClasses, _ := button.Attr("class")
		isNextButton = strings.Contains(buttonClasses, "BtnNext")

		// fmt.Println("Check Event : ", isNextButton, strings.TrimSuffix(strings.TrimPrefix(onclickEvent, "location.href='"), "'"))
	})
	defer res.Body.Close() // 어디에 배치하던 결국 마지막에 실행
	if !isNextButton {
		return buttons.Length()
	}
	return getTotalPageCount(buttons.Length()+1, searchKeyword) + buttons.Length()
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status:", res.StatusCode)
	}
}
