package security

type User interace {
	Name() string
	Id() string
	Groups() []string
}


type UserGetOperationBase struct {

}

type UserSetOperationBase struct {

}
