module local.packages/chap_02

go 1.17

replace local.packages/chap_02/popcount => ./chap_02/popcount

replace local.packages/chap_02/tempconv => ./chap_02/tempconv

require (
	gopl.io v0.0.0-20211004154805-1ae3ec64947b // indirect
	local.packages/chap_02/popcount v0.0.0-00010101000000-000000000000 // indirect
	local.packages/chap_02/tempconv v0.0.0-00010101000000-000000000000 // indirect
)
