/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcombn.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 18:10:52 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/30 17:54:33 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintCombN(n int) {
	num_set := []rune{}
	PrintHelper(num_set, 0, 0, n)
	ft.PrintRune('\n')
}

func PrintHelper(num_set []rune, selected int, start rune, n int) {
	if selected == n {
		PrintDelimiter(num_set, n)
		PrintList(num_set, n)
	} else {
		for v:='0' + start; v <= '9'; v++ {
			num_set = append(num_set, v)
			PrintHelper(num_set, selected + 1, v - '0' + 1, n)
			num_set = num_set[:len(num_set) - 1]
		}
	}
}

func PrintList(targ_list []rune, n int) {
	for index := 0; index < n; index++ {
		ft.PrintRune(targ_list[index])
	}
}

func PrintDelimiter(num_set []rune, n int) {
	not_min_flg := 0;
	var index_rune rune = '0';
	for index := 0; index < n; index++ {
		if num_set[index] != index_rune {
			not_min_flg = 1
		}
		index_rune++;
	}
	if not_min_flg == 1 {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}
