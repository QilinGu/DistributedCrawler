/*
 * 下载器(Downloader)
 * 用于下载网页内容, 并将网页内容返回给蜘蛛(Scrapy下载器是建立在twisted这个高效的异步模型上的)
 */
 package scrawler

 import (
   "fmt"
   "regexp"
   "strconv"
   "strings"
   "os"
   "bufio"
   "path/filepath"
 )