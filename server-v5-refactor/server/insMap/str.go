package insMap

import (
	"server-v5-refactor-server/static/structure"
	"server-v5-refactor-server/valueMap"
	"strconv"
)

func set() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "set",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			// get key and value from instruction.Values
			if len(instruction.Values) != 2 {
				return []string{"ERR", "wrong number of arguments for 'set' command"}
			} else {
				valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
					Type: structure.ObjString,
					Ptr:  instruction.Values[1],
				})
				return []string{"OK"}
			}
		},
	}
}

func get() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "get",
			BelongTo:       structure.ObjString,
			WillChangeData: false,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			// get key from instruction.Values
			if len(instruction.Values) != 1 {
				return []string{"ERR", "wrong number of arguments for 'get' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil {
					return []string{"nil"}
				} else {
					return []string{value.Ptr.(string)}
				}
			}
		},
	}
}

func del() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "del",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			// get key from instruction.Values
			if len(instruction.Values) != 1 {
				return []string{"ERR", "wrong number of arguments for 'del' command"}
			} else {
				valueMap.DelValueByKey(instruction.Values[0])
				return []string{"OK"}
			}
		},
	}
}

func strlen() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "strlen",
			BelongTo:       structure.ObjString,
			WillChangeData: false,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			// get key from instruction.Values
			if len(instruction.Values) != 1 {
				return []string{"ERR", "wrong number of arguments for 'strlen' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil {
					return []string{"0"}
				} else {
					return []string{strconv.Itoa(len(value.Ptr.(string)))}
				}
			}
		},
	}
}

func appendStr() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "append",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			// get key and value from instruction.Values
			if len(instruction.Values) != 2 {
				return []string{"ERR", "wrong number of arguments for 'append' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil {
					valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
						Type: structure.ObjString,
						Ptr:  instruction.Values[1],
					})
					return []string{strconv.Itoa(len(instruction.Values[1]))}
				} else {
					value.Ptr = value.Ptr.(string) + instruction.Values[1]
					return []string{strconv.Itoa(len(value.Ptr.(string)))}
				}
			}
		},
	}
}

func incr() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "incr",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			if len(instruction.Values) != 1 {
				return []string{"ERR", "wrong number of arguments for 'incr' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil { // if the key does not exist, create it with value 1
					valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
						Type: structure.ObjString,
						Ptr:  "1",
					})
					return []string{"1"}
				} else {
					// check if the value is a number
					num, err := strconv.Atoi(value.Ptr.(string))
					if err != nil {
						return []string{"ERR", "value is not an integer or out of range"}
					} else {
						value.Ptr = strconv.Itoa(num + 1)
						return []string{value.Ptr.(string)}
					}
				}
			}
		},
	}
}

func decr() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "decr",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			if len(instruction.Values) != 1 {
				return []string{"ERR", "wrong number of arguments for 'decr' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil { // if the key does not exist, create it with value -1
					valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
						Type: structure.ObjString,
						Ptr:  "-1",
					})
					return []string{"-1"}
				} else {
					// check if the value is a number
					num, err := strconv.Atoi(value.Ptr.(string))
					if err != nil {
						return []string{"ERR", "value is not an integer or out of range"}
					} else {
						value.Ptr = strconv.Itoa(num - 1)
						return []string{value.Ptr.(string)}
					}
				}
			}
		},
	}
}

func incrby() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "incrby",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			if len(instruction.Values) != 2 {
				return []string{"ERR", "wrong number of arguments for 'incrby' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil { // if the key does not exist, create it with value the second argument
					valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
						Type: structure.ObjString,
						Ptr:  instruction.Values[1],
					})
					return []string{instruction.Values[1]}
				} else {
					// check if the value is a number
					num, err := strconv.Atoi(value.Ptr.(string))
					if err != nil {
						return []string{"ERR", "value is not an integer or out of range"}
					} else {
						incr, err := strconv.Atoi(instruction.Values[1])
						if err != nil {
							return []string{"ERR", "step is not an integer or out of range"}
						} else {
							value.Ptr = strconv.Itoa(num + incr)
							return []string{value.Ptr.(string)}
						}
					}
				}
			}
		},
	}
}

func decrby() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "decrby",
			BelongTo:       structure.ObjString,
			WillChangeData: true,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) []string {
			if len(instruction.Values) != 2 {
				return []string{"ERR", "wrong number of arguments for 'decrby' command"}
			} else {
				value := valueMap.GetValueByKey(instruction.Values[0])
				if value == nil { // if the key does not exist, create it with value the second argument
					valueMap.SetValueByKey(instruction.Values[0], &structure.RedisObject{
						Type: structure.ObjString,
						Ptr:  instruction.Values[1],
					})
					return []string{instruction.Values[1]}
				} else {
					// check if the value is a number
					num, err := strconv.Atoi(value.Ptr.(string))
					if err != nil {
						return []string{"ERR", "value is not an integer or out of range"}
					} else {
						decr, err := strconv.Atoi(instruction.Values[1])
						if err != nil {
							return []string{"ERR", "step is not an integer or out of range"}
						} else {
							value.Ptr = strconv.Itoa(num - decr)
							return []string{value.Ptr.(string)}
						}
					}
				}
			}
		},
	}
}
