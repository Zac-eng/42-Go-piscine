/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 17:47:13 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 17:54:25 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func PrintComb() {
	for i:='0'; i<='9'; i++ {
		for j:=i + 1; j<='9'; j++ {
			for k:=j + 1; k<='9'; k++ {
				if i != '0' || j != '1' || k != '2' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
				ft.PrintRune(i)
				ft.PrintRune(j)
				ft.PrintRune(k)
			}
		}
	}
	ft.PrintRune('\n')
}
