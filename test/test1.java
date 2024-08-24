use strings;

out "第一句";
out 114+514;
out strings::merge("hello" , "world");

fun int_add : a int b int : int{
    ret a+b;
}

out int_add(10, 20);