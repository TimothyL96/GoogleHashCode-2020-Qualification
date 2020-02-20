cd ../

7z a qualification/output_best/source.zip stdlib/*
7z a qualification/output_last/source.zip stdlib/*
cd qualification

7z a output_best/source.zip main.go algorithm.go library.go read.go write.go
7z a output_last/source.zip main.go algorithm.go library.go read.go write.go