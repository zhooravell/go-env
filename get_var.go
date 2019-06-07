package env

// A set of functions for receiving data
// from the environment in the required format.
// It is wrapper over standard function os.Getenv(key).

import (
	"os"
	"strconv"
)

//GetVar return environment variable or default value
func GetVar(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}

//GetBytesVar return bytes environment variable or default value
func GetBytesVar(key string, fallback []byte) []byte {
	if v := os.Getenv(key); v != "" {
		return []byte(v)

	}

	return fallback
}

//GetBoolVar return bool environment variable or default value
func GetBoolVar(key string, fallback bool) (bool, error) {
	if v := os.Getenv(key); v != "" {
		return strconv.ParseBool(v)

	}

	return fallback, nil
}

//GeIntVar return int environment variable or default value
func GeIntVar(key string, fallback int) (int, error) {
	if v := os.Getenv(key); v != "" {
		return strconv.Atoi(v)
	}

	return fallback, nil
}

//GeInt8Var return int8 environment variable or default value
func GeInt8Var(key string, fallback int8) (int8, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseInt(v, 10, 8)

		if err != nil {
			return 0, err
		}

		return int8(i), nil
	}

	return fallback, nil
}

//GeInt16Var return int16 environment variable or default value
func GeInt16Var(key string, fallback int16) (int16, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseInt(v, 10, 16)

		if err != nil {
			return 0, err
		}

		return int16(i), nil
	}

	return fallback, nil
}

//GeInt32Var return int32 environment variable or default value
func GeInt32Var(key string, fallback int32) (int32, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseInt(v, 10, 32)

		if err != nil {
			return 0, err
		}

		return int32(i), nil
	}

	return fallback, nil
}

//GeInt64Var return int64 environment variable or default value
func GeInt64Var(key string, fallback int64) (int64, error) {
	if v := os.Getenv(key); v != "" {
		return strconv.ParseInt(v, 10, 64)
	}

	return fallback, nil
}

//GeUintVar return uint environment variable or default value
func GeUintVar(key string, fallback uint) (uint, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseUint(v, 10, 0)

		if err != nil {
			return 0, err
		}

		return uint(i), nil
	}

	return fallback, nil
}

//GeUint8Var return uint8 environment variable or default value
func GeUint8Var(key string, fallback uint8) (uint8, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseUint(v, 10, 8)

		if err != nil {
			return 0, err
		}

		return uint8(i), nil
	}

	return fallback, nil
}

//GeUint16Var return uint16 environment variable or default value
func GeUint16Var(key string, fallback uint16) (uint16, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseUint(v, 10, 16)

		if err != nil {
			return 0, err
		}

		return uint16(i), nil
	}

	return fallback, nil
}

//GeUint32Var return uint32 environment variable or default value
func GeUint32Var(key string, fallback uint32) (uint32, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseUint(v, 10, 32)

		if err != nil {
			return 0, err
		}

		return uint32(i), nil
	}

	return fallback, nil
}

//GeUint64Var return uint64 environment variable or default value
func GeUint64Var(key string, fallback uint64) (uint64, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseUint(v, 10, 64)

		if err != nil {
			return 0, err
		}

		return uint64(i), nil
	}

	return fallback, nil
}

//GetFloat64Var return float64 environment variable or default value
func GetFloat64Var(key string, fallback float64) (float64, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return 0, err
		}

		return float64(i), nil
	}

	return fallback, nil
}

//GetFloat32Var return float32 environment variable or default value
func GetFloat32Var(key string, fallback float32) (float32, error) {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return 0, err
		}

		return float32(i), nil
	}

	return fallback, nil
}
