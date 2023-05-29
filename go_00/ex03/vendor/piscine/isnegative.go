/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isnegative.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 17:39:50 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 17:44:55 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func IsNegative(nb int) {
	if nb < 0 {
		ft.PrintRune('T')
		ft.PrintRune('\n')
	} else {
		ft.PrintRune('F')
		ft.PrintRune('\n')
	}
}
