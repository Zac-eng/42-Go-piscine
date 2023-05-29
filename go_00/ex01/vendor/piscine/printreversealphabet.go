/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printreversealphabet.go                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 16:16:56 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 16:25:48 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func PrintReverseAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		ft.PrintRune(i);
	}
	ft.PrintRune('\n')
}
