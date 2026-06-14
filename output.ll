source_filename = "TestIs.mod"

@p3 = global { i32, i32, i32 } zeroinitializer
@result = global i32 zeroinitializer

define void @Check() {
entry:
	%0 = alloca {}
	%1 = getelementptr { i32, i32, i32 }, { i32, i32, i32 }* @p3, i32 0, i32 2
	store i32 5, i32* %1
	br i1 true, label %ifthen_0_3, label %ifelse_2

ifmerge_1:
	ret void

ifelse_2:
	store i32 1, i32* @result
	br label %ifmerge_1

ifthen_0_3:
	store i32 0, i32* @result
	br label %ifmerge_1
}

define void @__init_TestIs() {
entry:
	call void @Check()
	ret void
}

define i32 @main() {
entry:
	call void @__init_TestIs()
	%0 = load i32, i32* @result
	ret i32 %0
}
