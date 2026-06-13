source_filename = "TestFibonacci.mod"

@result = global i32 zeroinitializer

define i32 @Fib(i32 %0) {
entry:
	%1 = alloca { i32 }
	%2 = alloca i32
	%3 = getelementptr { i32 }, { i32 }* %1, i32 0, i32 0
	store i32 %0, i32* %3
	%4 = getelementptr { i32 }, { i32 }* %1, i32 0, i32 0
	%5 = load i32, i32* %4
	%6 = icmp sle i32 %5, 1
	br i1 %6, label %ifthen_0_3, label %ifelse_2

ifmerge_1:
	%7 = load i32, i32* %2
	ret i32 %7

ifelse_2:
	%8 = getelementptr { i32 }, { i32 }* %1, i32 0, i32 0
	%9 = load i32, i32* %8
	%10 = sub i32 %9, 1
	%11 = call i32 @Fib(i32 %10)
	%12 = getelementptr { i32 }, { i32 }* %1, i32 0, i32 0
	%13 = load i32, i32* %12
	%14 = sub i32 %13, 2
	%15 = call i32 @Fib(i32 %14)
	%16 = add i32 %11, %15
	store i32 %16, i32* %2
	br label %ifmerge_1

ifthen_0_3:
	%17 = getelementptr { i32 }, { i32 }* %1, i32 0, i32 0
	%18 = load i32, i32* %17
	store i32 %18, i32* %2
	br label %ifmerge_1
}

define void @__init_TestFibonacci() {
entry:
	%0 = call i32 @Fib(i32 10)
	%1 = icmp eq i32 %0, 55
	br i1 %1, label %ifthen_0_6, label %ifelse_5

ifmerge_4:
	ret void

ifelse_5:
	store i32 1, i32* @result
	br label %ifmerge_4

ifthen_0_6:
	store i32 0, i32* @result
	br label %ifmerge_4
}

define i32 @main() {
entry:
	call void @__init_TestFibonacci()
	%0 = load i32, i32* @result
	ret i32 %0
}
