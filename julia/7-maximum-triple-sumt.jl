function maxtrisum(arr)
	set_arr = []
	for i in eachindex(arr)
		if !(arr[i] in set_arr)
			append!(set_arr, arr[i])
		end
	end
	return sum(sort(arr, rev=true)[1:3])
end


println(maxtrisum([3,2,6,8,2,3]))


## -------------------------------------
function maxtrisum_2(arr)
	sum(sort(unique(arr))[end-2:end])
end


println(maxtrisum_2([3,2,6,8,2,3]))