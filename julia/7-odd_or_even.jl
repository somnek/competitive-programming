function oddoreven(array)
	return(sum(array) % 2 == 0 ? "even" : "odd")

end

println(oddoreven([0, 2, 3]))

# ------------------------------
# 				   arg      func   func  if true : else
oddoreven(array) = array |> sum |> isodd ? "odd" : "even"

oddoreven(array) = isodd(sum(array)) ? "odd" : "even"