function dutyfree(p, d, h )
    x = h/(p*d/100)
    return floor(Int, h/(p*d/100))
end

# solutions
dutyfree_1(p, d, h) = h ÷ (p * d /100)
dutyfree_2(p, d, h) = div(h, (p * d /100))


println(dutyfree(12, 50, 1000)) #166
