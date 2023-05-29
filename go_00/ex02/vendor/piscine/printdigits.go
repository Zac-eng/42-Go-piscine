/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printdigits.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 16:16:56 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 16:30:51 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func PrintDigits() {
	for i := '0'; i <= '9'; i++ {
		ft.PrintRune(i);
	}
	ft.PrintRune('\n')
}