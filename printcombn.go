package piscine

import "github.com/01-edu/z01"

func PrintComb01() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
		if i == '9' {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	z01.PrintRune('\n')
}

func PrintComb02() {
	for i := '0'; i <= '8'; i++ {
		for j := i + 1; j <= '9'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if i == '8' {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func PrintComb03() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i == '7' {
					break
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb04() {
	for i := '0'; i <= '6'; i++ {
		for j := i + 1; j <= '7'; j++ {
			for k := j + 1; k <= '8'; k++ {
				for l := k + 1; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(l)
					if i == '6' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb05() {
	for i := '0'; i <= '5'; i++ {
		for j := i + 1; j <= '6'; j++ {
			for k := j + 1; k <= '7'; k++ {
				for l := k + 1; l <= '8'; l++ {
					for m := l + 1; m <= '9'; m++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(m)
						if i == '5' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb06() {
	for i := '0'; i <= '4'; i++ {
		for j := i + 1; j <= '5'; j++ {
			for k := j + 1; k <= '6'; k++ {
				for l := k + 1; l <= '7'; l++ {
					for m := l + 1; m <= '8'; m++ {
						for n := m + 1; n <= '9'; n++ {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(m)
							z01.PrintRune(n)
							if i == '4' {
								break
							}
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb07() {
	for i := '0'; i <= '3'; i++ {
		for j := i + 1; j <= '4'; j++ {
			for k := j + 1; k <= '5'; k++ {
				for l := k + 1; l <= '6'; l++ {
					for m := l + 1; m <= '7'; m++ {
						for n := m + 1; n <= '8'; n++ {
							for o := n + 1; o <= '9'; o++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(m)
								z01.PrintRune(n)
								z01.PrintRune(o)
								if i == '3' {
									break
								}
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb08() {
	for i := '0'; i <= '2'; i++ {
		for j := i + 1; j <= '3'; j++ {
			for k := j + 1; k <= '4'; k++ {
				for l := k + 1; l <= '5'; l++ {
					for m := l + 1; m <= '6'; m++ {
						for n := m + 1; n <= '7'; n++ {
							for o := n + 1; o <= '8'; o++ {
								for p := o + 1; p <= '9'; p++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(n)
									z01.PrintRune(o)
									z01.PrintRune(p)
									if i == '2' {
										break
									}
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb09() {
	for i := '0'; i <= '1'; i++ {
		for j := i + 1; j <= '2'; j++ {
			for k := j + 1; k <= '3'; k++ {
				for l := k + 1; l <= '4'; l++ {
					for m := l + 1; m <= '5'; m++ {
						for n := m + 1; n <= '6'; n++ {
							for o := n + 1; o <= '7'; o++ {
								for p := o + 1; p <= '8'; p++ {
									for q := p + 1; q <= '9'; q++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(o)
										z01.PrintRune(p)
										z01.PrintRune(q)
										if i == '1' {
											break
										}
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintCombN(n int) {
	if n == 1 {
		PrintComb01()
	}
	if n == 2 {
		PrintComb02()
	}
	if n == 3 {
		PrintComb03()
	}
	if n == 4 {
		PrintComb04()
	}
	if n == 5 {
		PrintComb05()
	}
	if n == 6 {
		PrintComb06()
	}
	if n == 7 {
		PrintComb07()
	}
	if n == 8 {
		PrintComb08()
	}
	if n == 9 {
		PrintComb09()
	}
}
