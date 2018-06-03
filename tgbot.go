package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	// "fmt"
	"log"
	"strconv"
	"os"
	_ "github.com/lib/pq"
	"database/sql"
	// s "strings"
)

var (
	BOT_TOKEN = "534156580:AAGkSrGWqJPt3ygzCLTNOW3l9NQL_bepJkU"
	tgbot *tgbotapi.BotAPI
	GROUP_ID int64 = -288538740
)

func AddNewUser(usr int64){ // добавление нового пользователя в базу
	
}

func SendInfoMessage(m string){ //рассылка сообщения всем пользователям базы данны
	
}

// класс трайд бота

type TradeBot struct {
	Name string
	Number int
	Prc *os.Process
	Online bool
}

func (tb * TradeBot) Init(name int) { 		// инициализация бота строкой-инменем
	var t TradeBot			       	// позволяет
	t.Name = "bot" + strconv.Itoa(name)     // использовать общий путь + bot.Name
	t.Number = name			        // для доступа к файлам бота
	t.Online = false
	return
}

func (tb * TradeBot) WaitFor() {                    // метод ожидания для отлова
	tb.Prc.Wait()									// несанкционированного вылета трайд бота
	if tb.Online {                               // на случай, если выполнение было прервано функцией Stop
									                // чтобы не выводилось несколько раз сообщение
		msg := tgbotapi.NewMessage(GROUP_ID,tb.Name + "POWER OFF")
		msg.ParseMode = tgbotapi.ModeMarkdown
		tgbot.Send(msg)
	}
	tb.Online = false
}

func (tb * TradeBot) Go(args string){ // запуск бота
	if tb.Online {
		msg := tgbotapi.NewMessage(GROUP_ID,tb.Name + "ALREADY IS ONLINE")
		msg.ParseMode = tgbotapi.ModeMarkdown
		tgbot.Send(msg)
		return
	}
	var tb2 TradeBot
	tb2 = tb2.Init(tb.Number)
	tb2.Online = true
	var attr os.ProcAttr
	file,_ := os.Create(FILE + tb.Name + "/err_out_bot")
	str := []string{"main.js",args}
	attr.Files=[]*os.File{nil,nil,file}
	tb2.Prc,_ = os.StartProcess(FILE + tb.Name + "/main.js",str,&attr)
	go tb2.WaitFor() // функция запускается параллельно выполнению основного кода
	  		         // и ожидает завершения работы трейд бота

	if !tb.Online {
		return tb.Name + " ALREADY IS OFFLINE"
	}
	tb.Prc.Kill()
	tb.Online = false
	return tb.Name + " FORSED POWER OFF"
}

func (tb TradeBot) Errors() string {
	file, err := os.Open(FILE + tb.Name + "/errors")
	if(err != nil){
		os.Create(FILE + tb.Name + "/errors")
		file, err = os.Open(FILE + tb.Name + "/errors")
	}
	stat,_ := file.Stat()
	s := make([]byte, stat.Size())
	file.Read(s)
	return string(s)
}

func (tb TradeBot) Part() string {
	file, err := os.Open(FILE+tb.Name + "/part")
	if(err != nil){
		os.Create(FILE+tb.Name+"/part")
	}
	stat,_ := file.Stat()
	s := make([]byte, stat.Size())
	file.Read(s)
	return string(s)
}
//---------------------------------

func FileError(bots []*TradeBot){
	var botErr []string
	for _, i := bots {
		botErr = append(botErr,bots[i].Errors)
	}
	for {
		for i,_ := range bots{
			file, err := os.Open(FILE + tb.Name + "/errors")
			if(err != nil){
				os.Create(FILE + tb.Name + "/errors")
				file, err = os.Open(FILE + tb.Name + "/errors")
			}
			stat,_ := file.Stat()
			s := make([]byte, stat.Size())
			file.Read(s)
			if botErr[i] != s{
				SendErrorMessage(s)
				botErr[i] = s
			}
		}
	}
}
func main() {
	tgbot, err := tgbotapi.NewBotAPI(BOT_TOKEN)
	if err != nil {
	    log.Panic(err)
	}
	tgbot.Debug = true
	log.Printf("Authorized on account %s", tgbot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := tgbot.GetUpdatesChan(u)
	if err != nil {
	    log.Panic(err)
	}



	// В канал updates будут приходить все новые сообщения.
	for update := range updates {
		reply := ""

		cmd := update.Message
		log.Println(cmd.Command())
		switch cmd.Command(){
		case "hello":
			if cmd.CommandArguments() == "1337" {
				reply = strconv.FormatInt(update.Message.Chat.ID,10)
				AddNewUser(update.Message.Chat.ID)
			}
		case "send":
			SendInfoMessage("Hello, MyDaK!")
		case "info":


		}
	    msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	    tgbot.Send(msg)
	}
}