/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 17:57:04 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 18:07:25 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb2() {
	for i:='0'; i<='9'; i++ {
		for j:='0'; j<='9'; j++ {
			for k:= i; k<='9'; k++ {
				for l:=j + 1; l<='9'; l++ {
					if i != '0' || j != '0' || k != '0' || l != '1' {
						ft.PrintRune(',')
						ft.PrintRune(' ')
					}
						ft.PrintRune(i)
						ft.PrintRune(j)
						ft.PrintRune(' ')
						ft.PrintRune(k)
						ft.PrintRune(l)
				}
			}
		}
	}
	ft.PrintRune('\n')
}
