package main

import (
	"encoding/json"
	"errors"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// transaction will implement the processes
type SBITransaction struct {
}

type transEvent struct{
EventNo				string  `json:"eventNo"`
TransRefNo   		string  `json:"transRefNo"`
UserId				string 	`json:"userId"`
IpAdd   			string  `json:"ipAdd"`
EventDateTime   	string  `json:"eventDateTime"`
BCDateTime   		string  `json:"bcDateTime"`		// newly added by rathika
EventDesc   		string  `json:"eventDesc"`
Trans_branch		string  `json:"trans_branch"`
}

type sbiTransactions struct {

	TransRefNo   		string  `json:"ref_no"`
	RemAccNo     		string  `json:"rem_accno"`
	RemAmtINR    		float32 `json:"rem_amtinr"`
	BenAccNo	 		string  `json:"ben_accno"`		
	EventDesc    		string  `json:"event_desc"`	// have to delete later
	MakerUserID  		string  `json:"maker_id"`
	MakerIPAddr  		string  `json:"maker_ipaddr"`
	MakerDate    		string  `json:"maker_date"`
	AmlRemStatus    	string  `json:"aml_rem_status"`
    AmlBenStatus    	string  `json:"aml_ben_status"`
	OfacRemStatus   	string  `json:"ofac_rem_status"`
    OfacBenStatus   	string  `json:"ofac_ben_status"`
	RbiRemStatus   		string  `json:"rbi_rem_status"`
    RbiBenStatus   		string  `json:"rbi_ben_status"`
	Trans_init_branch	string  `json:"trans_init_branch"`
	Maker_branch		string  `json:"maker_branch"`

	L1UserID     		string  `json:"l1_userid"`
	L1IPAddr     		string  `json:"l1_ipaddr"`
	L1Status	 		string  `json:"l1_status"`
    L1AmlStatus     	string  `json:"l1_aml_status"`
    L1OfacStatus    	string  `json:"l1_ofac_status"`
    L1RbiStatus     	string  `json:"l1_rbi_status"`

	L2UserID    		string  `json:"l2_userid"`
	L2IPAddr     		string  `json:"l2_ipaddr"`
	L2Status	 		string  `json:"l2_status"`
	
	FinalcleStatus 	 	string  `json:"finacle_status"`
	
	TCSBancsStatus 		string  `json:"tcs_bancsstatus"`

	PSGStatus    		string  `json:"psg_status"`

	EventFlag			string  `json:"event_flag"`

    Key                 string  `json:"key"`

    //event data starts here
    DataEntryEventDate  string  `json:"data_entry_event_date"`
    DataEntryEventDesc  string  `json:"data_entry_event_desc"`

    AMLRemEventDate     string  `json:"aml_rem_event_date"`
    AMLRemEventDesc     string  `json:"aml_rem_event_desc"`

    AMLBenEventDate     string  `json:"aml_ben_event_date"`
    AMLBenEventDesc     string  `json:"aml_ben_event_desc"`

    OFACRemEventDate    string  `json:"ofac_rem_event_date"`
    OFACRemEventDesc    string  `json:"ofac_rem_event_desc"`

    OFACBenEventDate    string  `json:"ofac_ben_event_date"`
    OFACBenEventDesc    string  `json:"ofac_ben_event_desc"`

    RBIRemEventDate     string  `json:"rbi_rem_event_date"`
    RBIRemEventDesc     string  `json:"rbi_rem_event_desc"`

    RBIBenEventDate     string  `json:"rbi_ben_event_date"`
    RBIBenEventDesc     string  `json:"rbi_ben_event_desc"`

    L1EventDate         string  `json:"l1_event_date"`
    L1EventDesc         string  `json:"l1_event_desc"`

    L1AMLEventDate      string  `json:"l1_aml_event_date"`
    L1AMLEventDesc      string  `json:"l1_aml_event_desc"`

    L1OFACEventDate     string  `json:"l1_ofac_event_date"`
    L1OFACEventDesc     string  `json:"l1_ofac_event_desc"`

    L1RBIEventDate      string  `json:"l1_rbi_event_date"`
    L1RBIEventDesc      string  `json:"l1_rbi_event_desc"`

    L2EventDate         string  `json:"l2_event_date"`
    L2EventDesc         string  `json:"l2_event_desc"`

    FinacleEventDate    string  `json:"finacle_event_date"`
    FinacleEventDesc    string  `json:"finacle_event_desc"`

    TCSEventDate        string  `json:"tcs_event_date"`
    TCSEventDesc        string  `json:"tcs_event_desc"`

    PSGEventDate        string  `json:"psg_event_date"`
    PSGEventDesc        string  `json:"psg_event_desc"`

    //event data ends here
}

//Global declaration of maps
var sbi_trans_map map[string]sbiTransactions

//Invoke methods starts here 

//func CreateTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
func (t *SBITransaction) CreateTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var trans_obj sbiTransactions	
	var err error

	fmt.Println("Entering createTransaction")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		//return nil, errors.New("Expected atleast one arguments for initiate Transaction")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : ",args[0]," .. Args [1] is : ",args[1],"\n")

	//unmarshal transaction initiation data from UI to "sbiTransactions" struct
	err = json.Unmarshal([]byte(args[1]), &trans_obj)
	if err != nil {
		fmt.Printf("\nUnable to unmarshal createTransaction input transaction initiation : ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}
	fmt.Println("\nTransactionInitiation object refno variable value is : ",trans_obj.TransRefNo);

	// saving transactionInitiation and maker into map
	err = GetSBITransactionMap(stub)	

	//put transaction initiation data and maker data into map
    trans_obj.EventFlag = trans_obj.DataEntryEventDesc
	sbi_trans_map[trans_obj.TransRefNo] = trans_obj	

	err = SetSBITransactionMap(stub)	

	fmt.Printf("\ntransaction initiation map : ", sbi_trans_map)	
	fmt.Println("\nTransaction initiation Successfully saved")	
	
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}

	//return nil, nil
    return shim.Success(nil)
}

//func UpdateAML_OFAC_RBI(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
func (t *SBITransaction) UpdateAML_OFAC_RBI(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var trans_obj sbiTransactions	
    var bc_trans_obj sbiTransactions	
	var err error

	fmt.Println("Entering UpdateAML_OFAC_RBI")

	if (len(args) < 1) {
		fmt.Println("Invalid number of args")
		//return nil, errors.New("Expected atleast one arguments for initiate Transaction")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : ",args[0]," .. Args [1] is : ",args[1],"\n")

	//unmarshal transaction initiation data from UI to "sbiTransactions" struct
	err = json.Unmarshal([]byte(args[1]), &trans_obj)
	if err != nil {
		fmt.Printf("\nUnable to unmarshal createTransaction input transaction initiation : ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}
	fmt.Println("\n refno variable value is : ",trans_obj.TransRefNo);

	// saving transactionInitiation and maker into map
	err = GetSBITransactionMap(stub)


    // couch db query starts here

    queryString := fmt.Sprintf("{\"selector\":{\"TransRefNo\":{\"$eq\": trans_obj.TransRefNo}}}")
    queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
    fmt.Printf("\n queryResults for couch db : ",queryResults)
	//return shim.Success(queryResults)

    // couch db query ends here

    if err != nil {
		fmt.Printf("\nUnable to get map from blockchain : ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}

    fmt.Printf("\nBefore UpdateAML_OFAC_RBI map : %s ", sbi_trans_map)	
    fmt.Printf("\n Key  : %s ",trans_obj.Key)

    bc_trans_obj = sbi_trans_map[trans_obj.TransRefNo]

	//put transaction initiation data and maker data into map
   
    if(trans_obj.Key=="AML_Rem"){

         bc_trans_obj.EventFlag       = trans_obj.AMLRemEventDesc
         bc_trans_obj.AMLRemEventDesc = trans_obj.AMLRemEventDesc
         bc_trans_obj.AMLRemEventDate = trans_obj.AMLRemEventDate

    } else if(trans_obj.Key=="AML_Ben"){

         bc_trans_obj.EventFlag       = trans_obj.AMLBenEventDesc
         bc_trans_obj.AMLBenEventDesc = trans_obj.AMLBenEventDesc
         bc_trans_obj.AMLBenEventDate = trans_obj.AMLBenEventDate

    } else if(trans_obj.Key=="OFAC_Rem"){

         bc_trans_obj.EventFlag       = trans_obj.OFACRemEventDesc
         bc_trans_obj.OFACRemEventDesc = trans_obj.OFACRemEventDesc
         bc_trans_obj.OFACRemEventDate = trans_obj.OFACRemEventDate

    } else if(trans_obj.Key=="OFAC_Ben"){

         bc_trans_obj.EventFlag       = trans_obj.OFACBenEventDesc
         bc_trans_obj.OFACBenEventDesc = trans_obj.OFACBenEventDesc
         bc_trans_obj.OFACBenEventDate = trans_obj.OFACBenEventDate

    } else if(trans_obj.Key=="RBI_Rem"){

         bc_trans_obj.EventFlag       = trans_obj.RBIRemEventDesc
         bc_trans_obj.RBIRemEventDesc = trans_obj.RBIRemEventDesc
         bc_trans_obj.RBIRemEventDate = trans_obj.RBIRemEventDate

    } else if(trans_obj.Key=="RBI_Ben"){

         bc_trans_obj.EventFlag       = trans_obj.RBIBenEventDesc
         bc_trans_obj.RBIBenEventDesc = trans_obj.RBIBenEventDesc
         bc_trans_obj.RBIBenEventDate = trans_obj.RBIBenEventDate

    }
    fmt.Printf("\nbc_trans_obj.RBIBenEventDesc : %s ", bc_trans_obj.RBIBenEventDesc)
    fmt.Printf("\ntrans_obj.RBIBenEventDesc : %s ", trans_obj.RBIBenEventDesc)	
    fmt.Printf("\n")
	sbi_trans_map[bc_trans_obj.TransRefNo] = bc_trans_obj	

	err = SetSBITransactionMap(stub)	

    err = GetSBITransactionMap(stub)
	fmt.Printf("\nAfter UpdateAML_OFAC_RBI map : ", sbi_trans_map)	
	fmt.Println("\nTransaction initiation Successfully saved")
	
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}

	//return nil, nil
    return shim.Success(nil)
}


//func UpdateL1AuthorizerDetails(stub shim.ChaincodeStubInterface, args1 []string) error {
func (t *SBITransaction) UpdateL1AuthorizerDetails(stub shim.ChaincodeStubInterface, args1 []string) pb.Response {
	
	var l1_obj sbiTransactions	
    var bc_l1_obj sbiTransactions	
	var err error

	fmt.Println("Entering UpdateL1AuthorizerDetails\n")

	if (len(args1) < 1) {
		fmt.Println("Invalid number of args\n")
		//return errors.New("Expected atleast one arguments for UpdateL1AuthorizerDetails")
		return shim.Error(err.Error())
	}

	//unmarshal l1Authorizer data from UI to "l1Auth" struct
	err = json.Unmarshal([]byte(args1[1]), &l1_obj)
	if err != nil {
		fmt.Printf("Unable to marshal  createTransaction input UpdateL1AuthorizerDetails : %s\n", err)
		//return nil
		return shim.Error(err.Error())
	}

	// saving l1Authorizer details into map
		err = GetSBITransactionMap(stub)

    if err != nil {
		fmt.Printf("\nUnable to get map from blockchain : ", err)
		//return nil, nil
		return shim.Error(err.Error())
	}

    fmt.Printf("\nBefore UpdateL1AuthorizerDetails map : %s ", sbi_trans_map)	

	bc_l1_obj = sbi_trans_map[l1_obj.TransRefNo] 

    bc_l1_obj.L1UserID =  l1_obj.L1UserID 
	bc_l1_obj.L1IPAddr = l1_obj.L1IPAddr
    bc_l1_obj.L1EventDate = l1_obj.L1EventDate
    bc_l1_obj.L1EventDesc  = l1_obj.L1EventDesc 
    bc_l1_obj.L1Status = l1_obj.L1Status
    bc_l1_obj.EventFlag = l1_obj.L1EventDesc 

    sbi_trans_map[l1_obj.TransRefNo] = bc_l1_obj

	err = SetSBITransactionMap(stub)
	
	fmt.Printf("L1Authorizer map : %v \n", sbi_trans_map)	
	fmt.Println("L1Authorizer details Successfully updated")	

	//return nil
    return shim.Success(nil)
}

func (t *SBITransaction) ListAllTransactionEvent(stub shim.ChaincodeStubInterface)  pb.Response{
	var err error
	var bytesRead []byte
	var sbi_event_list []sbiTransactions	

	fmt.Println("Entering AllTransactionEvents")

	err = GetSBITransactionMap(stub)

	if err != nil {
		fmt.Printf("Unable to read the list of AllTransactionEvents : %s\n", err)
		//return nil, err
		return shim.Error(err.Error())
	}

	for _, value := range sbi_trans_map {
		//fmt.Printf("Events Value : %v\n", value.transRefNo)
		sbi_event_list = append(sbi_event_list, value)
	}
// couch db query starts here
/*
    {
  "selector": {
    "_id": {
      "$gt": null
    }
  } }
*/

//queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"marble\",\"owner\":\"%s\"}}", owner)
queryString := fmt.Sprintf("{\"selector\":{\"TransRefNo\":{\"$gt\": null}}}")
queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
// couch db query ends here


	/*fmt.Printf("list of AllTransactionEvents : %v\n", sbi_event_list)
	bytesRead, err = json.Marshal(&sbi_event_list)
	fmt.Printf("list of AllTransactionEvents after Marshal : %v\n", bytesRead)
	if err != nil {
		fmt.Printf("Unable to return the list of AllTransactionEvents : %s\n", err)
		return nil, err
	}

	return bytesRead, nil*/
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err		
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err			
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func GetSBITransactionMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = stub.GetState("SBITransactionMap")
	if err != nil {
		fmt.Printf("\nFailed to get  Transaction for block chain :", err)
		return err
	}
	if len(bytesread) != 0 {
		fmt.Printf("\nSBITransactionMap map exists.\n")
		err = json.Unmarshal(bytesread, &sbi_trans_map)
		if err != nil {
			fmt.Printf("\nFailed to initialize  SBITransactionMap for block chain :", err)
			return err
		}
	} else {
		fmt.Printf("\nSBITransactionMap map does not exist. To be created.")
		sbi_trans_map = make(map[string]sbiTransactions)
		bytesread, err = json.Marshal(&sbi_trans_map)
		if err != nil {
			fmt.Printf("\nFailed to initialize  SBITransactionMap for block chain :", err)
			return err
		}
		err = stub.PutState("SBITransactionMap", bytesread)
		if err != nil {
			fmt.Printf("\nFailed to initialize  SBITransactionMap for block chain : ", err)
			return err
		}
	}
	//fmt.Printf("sbi_trans_map from GETSBITransactionMap : ",sbi_trans_map)
	return nil
}


//setTransactionInitiationMap
func SetSBITransactionMap(stub shim.ChaincodeStubInterface) error {
	var err error
	var bytesread []byte

	bytesread, err = json.Marshal(&sbi_trans_map)
	if err != nil {
		fmt.Printf("\nFailed to set the TransactionItemMap for block chain : ", err)
		return err
	}

	err = stub.PutState("SBITransactionMap", bytesread)
	if err != nil {
		fmt.Printf("\nFailed to set the SBITransactionMap ", err)
		return errors.New("Failed to set the SBITransactionMap")
	}

	return nil
}


// Init sets up the chaincode
//func (t *SBITransaction) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	func (t *SBITransaction) Init(stub shim.ChaincodeStubInterface) pb.Response{
	//fmt.Println("Inside INIT for test chaincode")
	//return nil, nil
	return shim.Success(nil)
}

// Query the chaincode
//func (t *SBITransaction) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	func (t *SBITransaction) Query(stub shim.ChaincodeStubInterface) pb.Response {

		function, args := stub.GetFunctionAndParameters()
		//fmt.Println("invoke is running " + function)
			
        if function == "ListAllTransactionEvent" {
		return t.ListAllTransactionEvent(stub,args)
	    }
        
	//return nil, nil
	return shim.Error("Received unknown function invocation")
}

// Invoke the function in the chaincode
//func (t *SBITransaction) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	func (t *SBITransaction) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

		function, args := stub.GetFunctionAndParameters()
		//fmt.Println("invoke is running " + function)

	if function == "CreateTransaction" {
		return t.CreateTransaction(stub,args)
	} else if function == "UpdateAML_OFAC_RBI" {
		return t.UpdateAML_OFAC_RBI(stub,args)
	}

	fmt.Println("Function not found")
	//return nil, nil
	return shim.Error("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(SBITransaction))
	if err != nil {
		//fmt.Println("Could not start SBITransaction")
	} else {
		//fmt.Println("SBITransaction successfully started")
	}

}

