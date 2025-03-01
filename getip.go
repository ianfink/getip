/* 
 * Copyright (C) 2025 Ian M. Fink.
 * All rights reserved.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
 * more details.
 * 
 * You should have received a copy of the GNU General Public License along
 * with this program. If not, please see: https://www.gnu.org/licenses
 *
 */

package main

/*
 * Imports
 */

import (
	"fmt"
	"strings"
	"io"
	"net/http"
)

/*
 * Global Consts
 */

const theURL = "http://checkip.dyndns.org/"

/*
 * Funcs
 */

/**********************************************************************/

/**
 * Name:	GetIpFromString
 *
 * @brief	Determine a string that includes an IP address from
 *			a small webpage
 *
 * @param	s	the string that includes the IP address from the small webpage
 *
 * @return	the string that includes the IP address
 *
 */

func GetIpFromString (s string) string {
	// Example s = "<html><head><title>Current IP Check</title>" +
	//	"</head><body>Current IP Address: 34.138.201.253</body></html>"

	s1 := strings.Split(s, "<body>")
	s2 := strings.Split(s1[1], "</body>")

	return s2[0]
} /* GetIpFromString */

/**********************************************************************/

func main() {
	theResponse, err := http.Get(theURL)
	if err != nil {
		panic(err)
	}
	defer theResponse.Body.Close()

	theHTML, err := io.ReadAll(theResponse.Body)
	if err != nil {
		panic(err)
	}

	tmp := string(theHTML[:])

	fmt.Println(GetIpFromString(strings.ToLower(tmp)))
} /* main */

/**********************************************************************/

/*
 * End of file: getip.go
 */

