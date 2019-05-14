# toml-encoder
To generate a TOML file from the existing go types such as structs, maps, etc.

# Requirement:
The idea is to generate a TOML file based on the existing GO types in a project. These GO types can be structs, maps, etc. These GO types can be pulled from another package or in the same package. For simplicity, in this project, we will use the STRUCT go type.

# Current Challenge: 
I have used the toml-encoder librarary to make this possible only if we convert all slices which are fileds in a STRUCT, to an array size of at least [1]. This will pull all sub-fields of that slice in the encoded TOML file. If on the other hand, these struct fields which are slices, are left as slices, it will only pull the name of the slice and NOT the related subfields of that slice (which can be another STRUCT or a map)
