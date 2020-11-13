automorphic2(n) = string(n*n)[end-sizeof(string(n)) + 1:end] == string(n) ? "Automorphic" : "Not!!"

# ------------------------
automorphic(n) = endswith(string(n^2), string(n)) ? "Automorphic" : "Not!!"


