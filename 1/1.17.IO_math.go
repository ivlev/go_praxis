	/*

	Записать по правилам изучаемого языка программирования следующие вы-
	ражения:

	а)√(x₁² + x₂²)
	б) x₁x₂ + x₁x₃ + x₂x₃
	в) v₀t + at²/2
	г) mv²/2 + mgh
	д) 1/R1 + 1/R2
	е) mgcos(a)
	ж) 2𝜋R
	з) b² - 4ac
	и) 𝛾*(m₁m₂)
	к) I²R
	л) absin(c)
	м) √(a²+b²-2abcos(c))
	н) (ad+bc)/ad
	о) √(1-sin²(x))
	п) 1/√(ax²+bx+c)
	р) (√(x+1)+√(x-1))/2√x
	с) |x|+|x+1|
	т) |1-|x||

	 */

	package main

	import ( // "math"
		"fmt"

		"math"
	)

	func main() {
		var x[4] float64
		var R[3] float64
		var m[3] float64
		var v[1] float64
//		var a, b, c, x, t, R, h, m, I float64 // И запомните: енвозможно определить с одним именем переменную и массив!!
 		var a, b, c, t, h, r, I, V, M, X float64
		const g  = 9.8 // ускорение свободного падения
		const 𝛾 = 0.577215664901532860606512090082402431042159 // Постоянная Эйлера — Маскерони
		const 𝜋 = 3.141592653589793238462643383279 // Число Пи - как-то задумался, а ведь оно есть в math

		for i:=0;i<=9;i++ {
			switch i {
			case 0: fmt.Print("a=")
				fmt.Scanf("%f",&a)
			case 1: fmt.Print("b=")
				fmt.Scanf("%f",&b)
			case 2: fmt.Print("c=")
				fmt.Scanf("%f",&c)
			case 3: fmt.Print("t=")
				fmt.Scanf("%f",&t)
			case 4: fmt.Print("h=")
				fmt.Scanf("%f",&h)
			case 5: fmt.Print("R=")
				fmt.Scanf("%f",&r)
			case 6: fmt.Print("I=")
				fmt.Scanf("%f",&I)
			case 7: fmt.Print("V=")
				fmt.Scanf("%f",&V)
			case 8: fmt.Print("m=")
				fmt.Scanf("%f",&M)
			case 9: fmt.Print("x=")
				fmt.Scanf("%f",&X)

			}
		}
		for i:=0; i<=3; i++ {
			switch i {
			case 0: fmt.Print("x₀=") // в связи с дурацкими условиям избыточно болтаются элементы массива с индексом "0"
				fmt.Scanf("%f",&x[i])
			case 1: fmt.Print("x₁=")
				fmt.Scanf("%f",&x[i])
			case 2: fmt.Print("x₂=")
				fmt.Scanf("%f",&x[i])
			case 3: fmt.Print("x₃=")
				fmt.Scanf("%f",&x[i])
			}
		}
		for i:=0; i<=2; i++ {
			switch i {
			case 0: fmt.Print("R₀=")
				fmt.Scanf("%f",&R[i])
			case 1: fmt.Print("R₁=")
				fmt.Scanf("%f",&R[i])
			case 2: fmt.Print("R₂=")
				fmt.Scanf("%f",&R[i])
			}

		}
		for i:=0; i<=3; i++ {
			switch i {
			case 0: fmt.Print("m₀=")
				fmt.Scanf("%f",&m[i])
			case 1: fmt.Print("m₁=")
				fmt.Scanf("%f",&m[i])
			case 2: fmt.Print("m₂=")
				fmt.Scanf("%f",&m[i])
			}

		}
		for i:=0; i<=0; i++ { // во избежиние конфликта имён массив нужно определять буквой в другим регистре
			switch i {
			case 0: fmt.Print("v₀=")
				fmt.Scanf("%f",&v[i])
			}

		}

		fmt.Printf("%s%f\n", "а) √(x₁² + x₂²)=", math.Sqrt(math.Pow(x[1],2)+math.Pow(x[2],2)))
		fmt.Printf("%s%f\n", "б) x₁x₂ + x₁x₃ + x₂x₃=", x[1]*x[2] + x[1]*x[3] + x[2]*x[3])
		fmt.Printf("%s%f\n", "в) v₀t + at²/2=", v[0]*t + (a*math.Pow(t, 2))/2)
		fmt.Printf("%s%f\n", "г) mv²/2 + mgh=", (M*math.Pow(t,2))/2 + M*g*h)
		fmt.Printf("%s%f\n", "д) 1/R1 + 1/R2=", 1/R[1] + 1/R[2])
		fmt.Printf("%s%f\n", "е) mgcos(a)=", M*g*math.Cos(a))
		fmt.Printf("%s%f\n", "ж) 2𝜋R=", 2*math.Pi*r)
		fmt.Printf("%s%f\n", "з) b² - 4ac=", math.Pow(b, 2) - 4*a*c)
		fmt.Printf("%s%f\n", "и) 𝛾*(m₁m₂)=", 𝛾*(m[1]*m[2]))
		fmt.Printf("%s%f\n", "к) I²R=", math.Pow(I, 2)*r)
		fmt.Printf("%s%f\n", "л) absin(c)=", a*b*math.Sin(c))
		fmt.Printf("%s%f\n", "м) √(a²+b²-2abcos(c))=", math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2) - 2*a*b*math.Cos(c)))
		fmt.Printf("%s%f\n", "н) (ad+bc)/ad=", (a*b+b*c)/a*b)
		fmt.Printf("%s%f\n", "о) √(1-sin²(x))=", math.Sqrt(1-math.Pow(math.Sin(X),2)))
		fmt.Printf("%s%f\n", "п) 1/√(ax²+bx+c)=", 1/math.Sqrt(a*math.Pow(X, 2)+b*X + c))
		fmt.Printf("%s%f\n", "р) (√(x+1)+√(x-1))/2√x=", (math.Sqrt(X+1) + math.Sqrt(X-1))/2*math.Sqrt(X))
		fmt.Printf("%s%f\n", "с) |x|+|x+1|=", math.Abs(X) + math.Abs(X+1))
		fmt.Printf("%s%f\n", "т) |1-|x||", math.Abs(1 - math.Abs(X)))
	}

	/*
	е) mgcos(a)
	ж) 2𝜋R
	з) b² - 4ac
	и) 𝛾*(m₁m₂)
	к) I²R
	л) absin(c)
	м) √(a²+b²-abcos(c))
	н) (ad+bc)/ad
	о) √(1-sin²(x))
	п) 1/√(ax²+bx+c)
	р) (√(x+1)+√(x-1))/2√x
	с) |x|+|x+1|
	т) |1-|x||

	*/