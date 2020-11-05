
function findshort(s)
	return minimum([sizeof(string(x)) for x in split(s)])
end


test_string  ="bitcoin take over the world maybe who knows perhaps"
println(findshort(test_string))

##----------------------------------------------------
#					map(func1, args for func1)
findshortx(s) = minimum(sizeof, split(s))

println(findshortx("asd dsdf sdf"))