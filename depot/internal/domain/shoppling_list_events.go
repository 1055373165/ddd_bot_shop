package domain

type ShoppingListCreated struct {
	ShoppingList *ShoppingList
}

func (ShoppingListCreated) EventName() string { return "depot.ShopplingListCreated" }

type ShoppingListCanceled struct {
	ShoppingList *ShoppingList
}

func (ShoppingListCanceled) EventName() string { return "depot.ShoppingListCanceled" }

type ShoppingListAssigned struct {
	ShoppingList *ShoppingList
	BotID        string
}

func (ShoppingListAssigned) EventName() string { return "depot.ShopplingListAssigned" }

type ShoppingListCompleted struct {
	ShoppingList *ShoppingList
}

func (ShoppingListCompleted) EventName() string { return "depot.ShoppingListCompleted" }
