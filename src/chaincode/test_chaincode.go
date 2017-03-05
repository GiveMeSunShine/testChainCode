package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"

	"encoding/json"
)

var myLogger = logging.MustGetLogger("user_chain")

type UserChaincode struct {
	Id  string `json:"id"`
	Name string `json:"name"`
	Identity_type string `json:"identity_type"`
	Identity_code string `json:"identity_code"`
	Account_name string `json:"account_name"`
	Card_branch_name string `json:"card_branch_name"`
	Card_id string `json:"card_id"`
	Cnaps_code string `json:"cnaps_code"`
	Phone_number string `json:"phone_number"`
	Home_address string `json:"home_address"`
	Education string `json:"education"`
	Company string `json:"company"`
	Job_number string `json:"job_number"`
	Job_address string `json:"job_address"`
	Job_title string `json:"job_title"`
	Certificate string `json:"certificate"`
	Experiments string `json:"experiments"`
	Rewards string `json:"rewards"`
	Personal_Valuation string `json:"personal_Valuation"`
	Interest string `json:"interest"`
}

func (t *UserChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *UserChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	user := UserChaincode{}
	user.Id = args[0]
	user.Name = args[1]
	user.Identity_type = args[2]
	user.Identity_code = args[3]
	user.Account_name = args[4]
	user.Card_branch_name = args[5]
	user.Card_id = args[6]
	user.Cnaps_code = args[7]
	user.Phone_number = args[8]
	user.Home_address = args[9]
	user.Education = args[10]
	user.Company = args[11]
	user.Job_number = args[12]
	user.Job_address = args[13]
	user.Job_title = args[14]
	user.Certificate = args[15]
	user.Experiments = args[16]
	user.Rewards = args[17]
	user.Personal_Valuation = args[18]
	user.Interest = args[19]
	myLogger.Info("UserChainCode is s%",user.Id)
	buf,err := json.Marshal(user)
	if err != nil {
		myLogger.Error("Failed Marsha")
		return nil, err
	}
	stub.PutState(user.Id,buf )
	return  nil,err
}

func (t *UserChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	user := UserChaincode{}
	buf,err:= stub.GetState(args[0])
	if err != nil {
		myLogger.Error("ERROR Query Failed !")
		return nil, err
	}
	if buf == nil {
		myLogger.Error("Query result is null !")
		return nil, err
	}
	json.Unmarshal(buf,user)
	return  buf,err
}

func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		myLogger.Debugf("Error starting UserChaincode: %s", err)
	}
}


