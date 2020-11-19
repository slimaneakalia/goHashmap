package hashMap

import "testing"

func TestGetValue(t *testing.T){
	listHead := &LinkedList{ value: "head 1" }
	if listHead.GetValue() != "head 1"{
		t.Errorf("GetValue not working, current value %v", listHead.GetValue())
	} else {
		t.Log("GetValue is working")
	}
}

func TestAddValue(t *testing.T){
	listHead := &LinkedList{ value: "head 1" }
	listHead = listHead.AddValue("head 2")

	if listHead.GetValue() != "head 2"{
		t.Errorf("AddValue not working, current value %v", listHead.GetValue())
	} else {
		t.Log("AddValue is working")
	}

}

func TestFindValue(t *testing.T){
	listHead := &LinkedList{ value: "head 1" }
	listHead = listHead.AddValue("head 2")
	listHead = listHead.AddValue("head 3")

	researchedValue := listHead.FindValue("head 3", nil)

	if researchedValue == nil {
		t.Error("FindValue is not working")
	} else {
		t.Log("FindValue is working")
	}
}

func TestPrint(t *testing.T){
	listHead := &LinkedList{ value: "head 1" }
	listHead = listHead.AddValue("head 2")
	listHead = listHead.AddValue("head 3")

	str := listHead.Print()

	if str != "head 3 - head 2 - head 1 - "{
		t.Errorf("Print is not working, current output: %v", str)
	}else {
		t.Log("Print is working")
	}
}