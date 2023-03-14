package main

import (
	"fmt"
	"go-course/reflection/common"
	"reflect"

	"github.com/go-playground/assert/v2"
)

func get_type() {
	typeI := reflect.TypeOf(1)       //通过TypeOf()得到Type类型
	typeS := reflect.TypeOf("hello") //通过TypeOf()得到Type类型
	fmt.Println(typeI)               //int
	fmt.Println(typeI.String())      //int
	fmt.Println(typeS)               //string
	fmt.Println(typeI.Kind() == reflect.Int)
	fmt.Println(typeS.Kind() == reflect.String)
	fmt.Println()

	typeUser := reflect.TypeOf(&common.User{}) //通过TypeOf()得到Type类型
	fmt.Println(typeUser)                      //*common.User
	fmt.Println(typeUser.Elem())               //common.User
	fmt.Println(typeUser.Name())               //空字符串
	fmt.Println(typeUser.Elem().Name())        //User，不带包名的类名称
	fmt.Println(typeUser.Kind())               //ptr
	fmt.Println(typeUser.Elem().Kind())        //struct
	fmt.Println(typeUser.Kind() == reflect.Ptr)
	fmt.Println(typeUser.Elem().Kind() == reflect.Struct)
	fmt.Println()

	typeUser2 := reflect.TypeOf(common.User{})
	fmt.Println(typeUser2)           //common.User
	fmt.Println(typeUser2.Name())    //User
	fmt.Println(typeUser2.Kind())    //struct
	fmt.Println(typeUser2.PkgPath()) //go-course/reflection/common，包路径
	fmt.Println(typeUser2.Size())    //48，占用内存大小
	// fmt.Println(typeUser2.Elem()) panic，非指针类型不能调用Elem()
	assert.IsEqual(typeUser.Elem(), typeUser2)
}

func get_field() {
	typeUser := reflect.TypeOf(common.User{}) //需要用struct的Type，不能用指针的Type
	fieldNum := typeUser.NumField()           //成员变量的个数
	for i := 0; i < fieldNum; i++ {
		field := typeUser.Field(i)
		fmt.Printf("%d %s offset %d anonymous %t type %s exported %t json tag %s\n", i,
			field.Name,            //变量名称
			field.Offset,          //相对于结构体首地址的内存偏移量，string类型会占据16个字节
			field.Anonymous,       //是否为匿名成员
			field.Type,            //数据类型，reflect.Type类型
			field.IsExported(),    //包外是否可见（即是否以大写字母开头）
			field.Tag.Get("json")) //获取成员变量后面``里面定义的tag
	}
	fmt.Println()

	//可以通过FieldByName获取Field
	if nameField, ok := typeUser.FieldByName("Name"); ok {
		fmt.Printf("Name is exported %t\n", nameField.IsExported())
	}
	//也可以根据FieldByIndex获取Field
	thirdField := typeUser.FieldByIndex([]int{2}) //参数是个slice，因为有struct嵌套的情况
	fmt.Printf("third field name %s\n", thirdField.Name)
}

func get_method() {
	typeUser := reflect.TypeOf(common.User{})
	methodNum := typeUser.NumMethod() //成员方法的个数。接收者为指针的方法【不】包含在内
	for i := 0; i < methodNum; i++ {
		method := typeUser.Method(i)
		fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}
	fmt.Println()

	typeUser2 := reflect.TypeOf(&common.User{})
	methodNum = typeUser2.NumMethod() //成员方法的个数。接收者为指针或值的方法都包含在内，也就是说值实现的方法指针也实现了（反之不成立）
	for i := 0; i < methodNum; i++ {
		method := typeUser2.Method(i)
		fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	}
}

func Add(a, b int) int {
	return a + b
}

func get_func() {
	typeFunc := reflect.TypeOf(Add) //获取函数类型
	fmt.Printf("is function type %t\n", typeFunc.Kind() == reflect.Func)
	argInNum := typeFunc.NumIn()   //输入参数的个数
	argOutNum := typeFunc.NumOut() //输出参数的个数
	for i := 0; i < argInNum; i++ {
		argTyp := typeFunc.In(i)
		fmt.Printf("第%d个输入参数的类型%s\n", i, argTyp)
	}
	for i := 0; i < argOutNum; i++ {
		argTyp := typeFunc.Out(i)
		fmt.Printf("第%d个输出参数的类型%s\n", i, argTyp)
	}
}

func get_struct_relation() {
	u := reflect.TypeOf(common.User{})
	t := reflect.TypeOf(common.Student{}) //Student内部嵌套了User
	u2 := reflect.TypeOf(common.User{})

	//false false
	fmt.Println(t.AssignableTo(u))  //t代表的类型是否可以赋值给u代表的类型
	fmt.Println(t.ConvertibleTo(u)) //t代表的类型是否可以转换成u代表的类型

	//java的反射可以获取继承关系，而go语言不支持继承，所以必须是相同的类型才能AssignableTo和ConvertibleTo
	//false false
	fmt.Println(u.AssignableTo(t))
	fmt.Println(u.ConvertibleTo(t))

	//true true
	fmt.Println(u.AssignableTo(u2))
	fmt.Println(u.ConvertibleTo(u2))
}

//判断某个struct是否实现了某个接口
func implements() {
	//通过reflect.TypeOf((*<interface>)(nil)).Elem()获得接口类型。因为People是个接口不能创建实例，所以把nil强制转为*common.People类型
	typeOfPeople := reflect.TypeOf((*common.People)(nil)).Elem()
	fmt.Printf("typeOfPeople kind is interface %t\n", typeOfPeople.Kind() == reflect.Interface)
	t1 := reflect.TypeOf(common.User{})
	t2 := reflect.TypeOf(&common.User{})
	//User的值类型实现了接口，则指针类型也实现了接口；但反过来不行(把Think的接收者改为*User试试)
	fmt.Printf("t1 implements People interface %t\n", t1.Implements(typeOfPeople))
	fmt.Printf("t2 implements People interface %t\n", t2.Implements(typeOfPeople))
	fmt.Println()
}

//Value和其他类型的互换
func value_other_conversion() {
	//原始类型转为Value
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("hello")
	userPtrValue := reflect.ValueOf(&common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65,
		Height: 1.68,
	})
	fmt.Println(iValue)       //1
	fmt.Println(sValue)       //hello
	fmt.Println(userPtrValue) //&{7 杰克逊  65 1.68}

	//Value转为Type
	iType := iValue.Type()
	sType := sValue.Type()
	userType := userPtrValue.Type()
	//在Type和相应Value上调用Kind()结果一样的
	fmt.Println(iType.Kind() == reflect.Int, iValue.Kind() == reflect.Int, iType.Kind() == iValue.Kind())                   //true true
	fmt.Println(sType.Kind() == reflect.String, sValue.Kind() == reflect.String, sType.Kind() == sValue.Kind())             //true true
	fmt.Println(userType.Kind() == reflect.Ptr, userPtrValue.Kind() == reflect.Ptr, userType.Kind() == userPtrValue.Kind()) //true true true

	//指针Value和非指针Value互相转换
	userValue := userPtrValue.Elem()                    //Elem() 指针Value转为非指针Value
	fmt.Println(userValue.Kind(), userPtrValue.Kind())  //struct ptr
	userPtrValue3 := userValue.Addr()                   //Addr() 非指针Value转为指针Value
	fmt.Println(userValue.Kind(), userPtrValue3.Kind()) //struct ptr

	//通过Interface()函数把Value转为interface{}，再从interface{}强制类型转换，转为原始数据类型
	//或者在Value上直接调用Int()、String()等一步到位
	fmt.Printf("origin value iValue is %d %d\n", iValue.Interface().(int), iValue.Int())
	fmt.Printf("origin value sValue is %s %s\n", sValue.Interface().(string), sValue.String())
	user := userValue.Interface().(common.User)
	fmt.Printf("id=%d name=%s weight=%.2f height=%.2f\n", user.Id, user.Name, user.Weight, user.Height)
	user2 := userPtrValue.Interface().(*common.User)
	fmt.Printf("id=%d name=%s weight=%.2f height=%.2f\n", user2.Id, user2.Name, user2.Weight, user2.Height)
}

func value_is_empty() {
	var i interface{} //接口没有指向具体的值
	v := reflect.ValueOf(i)
	fmt.Printf("v持有值 %t, type of v is Invalid %t\n", v.IsValid(), v.Kind() == reflect.Invalid)

	var user *common.User = nil
	v = reflect.ValueOf(user) //Value指向一个nil
	if v.IsValid() {
		fmt.Printf("v持有的值是nil %t\n", v.IsNil()) //调用IsNil()前先确保IsValid()，否则会panic
	}

	var u common.User //只声明，里面的值都是0值
	v = reflect.ValueOf(u)
	if v.IsValid() {
		fmt.Printf("v持有的值是对应类型的0值 %t\n", v.IsZero()) //调用IsZero()前先确保IsValid()，否则会panic
	}
}

//通过反射改变原始数据的值
func chane_value() {
	var i int = 10
	var s string = "hello"
	user := common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}

	valueI := reflect.ValueOf(&i) //由于go语言所有函数传的都是值，所以要想修改原来的值就需要传指针
	valueS := reflect.ValueOf(&s)
	valueUser := reflect.ValueOf(&user)
	valueI.Elem().SetInt(8)           //由于valueI对应的原始对象是指针，通过Elem()返回指针指向的对象
	valueS.Elem().SetString("golang") //不能在指针Value上调用Set系列函数
	// valueS.Elem().SetInt(8)//会panic
	valueUser.Elem().FieldByName("Weight").SetFloat(68.0) //FieldByName()通过Name返回类的成员变量。不能在指针Value上调用FieldByName
	fmt.Printf("i=%d, s=%s, user.Weight=%.1f\n", i, s, user.Weight)
	addrValue := valueUser.Elem().FieldByName("addr")
	if addrValue.CanSet() {
		addrValue.SetString("北京")
	} else {
		fmt.Println("addr是未导出成员，不可Set") //以小写字母开头的成员相当于是私有成员
	}
}

func change_slice() {
	users := make([]*common.User, 1, 5) //len=1，cap=5
	users[0] = &common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}

	sliceValue := reflect.ValueOf(&users) //准备通过Value修改users，所以传users的地址
	if sliceValue.Elem().Len() > 0 {      //取得slice的长度
		sliceValue.Elem().Index(0).Elem().FieldByName("Name").SetString("令狐一刀")
		// u0 := users[0]
		fmt.Printf("1st user name change to %s\n", users[0].Name)
	}

	sliceValue.Elem().SetCap(3) //新的cap必须位于原始的len到cap之间
	sliceValue.Elem().SetLen(2)
	//调用reflect.Value的Set()函数修改其底层指向的原始数据
	sliceValue.Elem().Index(1).Set(reflect.ValueOf(&common.User{
		Id:     8,
		Name:   "李达",
		Weight: 80,
		Height: 180,
	}))
	fmt.Printf("2nd user name %s\n", users[1].Name)
}

func change_map() {
	u1 := &common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}
	u2 := &common.User{
		Id:     8,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}
	userMap := make(map[int]*common.User, 5)
	userMap[u1.Id] = u1

	mapValue := reflect.ValueOf(&userMap)                                                         //准备通过Value修改userMap，所以传userMap的地址
	mapValue.Elem().SetMapIndex(reflect.ValueOf(u2.Id), reflect.ValueOf(u2))                      //SetMapIndex 往map里添加一个key-value对
	mapValue.Elem().MapIndex(reflect.ValueOf(u1.Id)).Elem().FieldByName("Name").SetString("令狐一刀") //MapIndex 根据Key取出对应的map
	for k, user := range userMap {
		fmt.Printf("key %d name %s\n", k, user.Name)
	}
}

//通过反射调用函数
func call_function() {
	valueFunc := reflect.ValueOf(Add) //函数也是一种数据类型
	typeFunc := reflect.TypeOf(Add)
	argNum := typeFunc.NumIn()            //函数输入参数的个数
	args := make([]reflect.Value, argNum) //准备函数的输入参数
	for i := 0; i < argNum; i++ {
		if typeFunc.In(i).Kind() == reflect.Int {
			args[i] = reflect.ValueOf(3) //给每一个参数都赋3
		}
	}
	sumValue := valueFunc.Call(args) //返回[]reflect.Value，因为go语言的函数返回可能是一个列表
	if typeFunc.Out(0).Kind() == reflect.Int {
		sum := sumValue[0].Interface().(int) //从Value转为原始数据类型
		fmt.Printf("sum=%d\n", sum)
	}
}

//通过反射调用方法
func call_method() {
	user := common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}
	valueUser := reflect.ValueOf(&user)              //必须传指针，因为BMI()在定义的时候它是指针的方法
	bmiMethod := valueUser.MethodByName("BMI")       //MethodByName()通过Name返回类的成员变量
	resultValue := bmiMethod.Call([]reflect.Value{}) //无参数时传一个空的切片
	result := resultValue[0].Interface().(float32)
	fmt.Printf("bmi=%.2f\n", result)

	//Think()在定义的时候用的不是指针，valueUser可以用指针也可以不用指针
	thinkMethod := valueUser.MethodByName("Think")
	thinkMethod.Call([]reflect.Value{})

	valueUser2 := reflect.ValueOf(user)
	thinkMethod = valueUser2.MethodByName("Think")
	thinkMethod.Call([]reflect.Value{})
}

func new_struct() {
	t := reflect.TypeOf(common.User{})
	value := reflect.New(t) //根据reflect.Type创建一个对象，得到该对象的指针，再根据指针提到reflect.Value
	value.Elem().FieldByName("Id").SetInt(10)
	value.Elem().FieldByName("Name").SetString("宋江")
	value.Elem().FieldByName("Weight").SetFloat(78.)
	value.Elem().FieldByName("Height").SetFloat(168.4)
	user := value.Interface().(*common.User) //把反射类型转成go原始数据类型
	fmt.Printf("id=%d name=%s weight=%.1f height=%.1f\n", user.Id, user.Name, user.Weight, user.Height)
}

func new_slice() {
	var slice []common.User
	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.MakeSlice(sliceType, 1, 3) //reflect.MakeMap、reflect.MakeSlice、reflect.MakeChan、reflect.MakeFunc
	sliceValue.Index(0).Set(reflect.ValueOf(common.User{
		Id:     8,
		Name:   "李达",
		Weight: 80,
		Height: 180,
	}))
	users := sliceValue.Interface().([]common.User)
	fmt.Printf("1st user name %s\n", users[0].Name)
}

func new_map() {
	var userMap map[int]*common.User
	mapType := reflect.TypeOf(userMap)
	// mapValue:=reflect.MakeMap(mapType)
	mapValue := reflect.MakeMapWithSize(mapType, 10) //reflect.MakeMap、reflect.MakeSlice、reflect.MakeChan、reflect.MakeFunc

	user := &common.User{
		Id:     7,
		Name:   "杰克逊",
		Weight: 65.5,
		Height: 1.68,
	}
	key := reflect.ValueOf(user.Id)
	mapValue.SetMapIndex(key, reflect.ValueOf(user))                    //SetMapIndex 往map里添加一个key-value对
	mapValue.MapIndex(key).Elem().FieldByName("Name").SetString("令狐一刀") //MapIndex 根据Key取出对应的map
	userMap = mapValue.Interface().(map[int]*common.User)
	fmt.Printf("user name %s %s\n", userMap[7].Name, user.Name)
}

func main() {
	get_type()
	// fmt.Println("===================")
	// get_field()
	// fmt.Println("===================")
	// get_method()
	// fmt.Println("===================")
	// get_func()
	// fmt.Println("===================")
	// get_struct_relation()
	// fmt.Println("===================")
	// implements()
	// fmt.Println("===================")

	// value_other_conversion()
	// fmt.Println("===================")
	// value_is_empty()
	// fmt.Println("===================")
	// chane_value()
	// fmt.Println("===================")
	// change_slice()
	// fmt.Println("===================")
	// change_map()
	// fmt.Println("===================")

	// call_function()
	// fmt.Println("===================")
	// call_method()
	// fmt.Println("===================")

	// new_struct()
	// fmt.Println("===================")
	// new_slice()
	// fmt.Println("===================")
	// new_map()
}

//go run reflection/basic/main.go
