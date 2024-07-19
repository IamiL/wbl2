package main

import "fmt"

type Bank struct {
	site     site
	db       db
	notifier notifier
}

type db struct {
}

func (db *db) ReduceBalance() {
	fmt.Println("уменьшаем баланс")
}

func (db *db) IncreaseBalance() {
	fmt.Println("увеличиваем баланс")
}

type site struct {
}

func (s *site) GetTransfer() {
	fmt.Println("получаем перевод с сайта")
}

func (s *site) CloseTransfer() {
	fmt.Println("отображаем выполненный перевод")
}

type notifier struct {
}

func (n *notifier) NotifyRecipient() {
	fmt.Println("уведомляем получателя о переводе")
}

func (n *notifier) NotifySender() {
	fmt.Println("уведомляем отправителя о переводе")
}

func (b *Bank) GetTransfer() {
	b.site.GetTransfer()
	b.notifier.NotifySender()
}

func (b *Bank) MakeTransfer() {
	b.db.ReduceBalance()
	b.db.IncreaseBalance()

	b.site.CloseTransfer()

	b.notifier.NotifyRecipient()
	b.notifier.NotifySender()
}

func main() {
	bank := &Bank{}

	bank.GetTransfer()
	bank.MakeTransfer()
}
