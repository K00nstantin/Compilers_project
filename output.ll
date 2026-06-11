source_filename = "TestReal.mod"

@r = global float zeroinitializer
@i = global i32 zeroinitializer
@ok = global i1 zeroinitializer
@result = global i32 zeroinitializer

define void @__init_TestReal() {
entry:
	store i32 2, i32* @i
	%0 = load i32, i32* @i
	%1 = sitofp i32 %0 to float
	%2 = fmul float 0x400921F9E0000000, %1
	store float %2, float* @r
	%3 = load float, float* @r
	%4 = fcmp ogt float %3, 6.0
	%5 = load float, float* @r
	%6 = fcmp olt float %5, 7.0
	%7 = and i1 %4, %6
	store i1 %7, i1* @ok
	%8 = load i1, i1* @ok
	br i1 %8, label %ifthen_0_3, label %ifelse_2

ifmerge_1:
	ret void

ifelse_2:
	store i32 1, i32* @result
	br label %ifmerge_1

ifthen_0_3:
	store i32 0, i32* @result
	br label %ifmerge_1
}

define i32 @main() {
entry:
	call void @__init_TestReal()
	%0 = load i32, i32* @result
	ret i32 %0
}
