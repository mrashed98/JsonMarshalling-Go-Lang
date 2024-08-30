module jsonMarshalling

go 1.22.5

replace unmarshall => ./unmarshall

replace jsonMarshalling/unmarshall => ./unmarshall

require (
	jsonMarshalling/marshall v0.0.0-00010101000000-000000000000
	jsonMarshalling/unmarshall v0.0.0-00010101000000-000000000000
)

replace jsonMarshalling/marshall => ./marshall
