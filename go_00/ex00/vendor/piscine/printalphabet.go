/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printalphabet.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: h.miyazaki <h.miyazaki@student.42.fr>      +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/29 16:16:56 by h.miyazaki        #+#    #+#             */
/*   Updated: 2023/05/29 16:21:31 by h.miyazaki       ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import (
	"ft"
)

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		ft.PrintRune(i);
	}
	ft.PrintRune('\n')
}
