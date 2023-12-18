package main 


import (
	"fmt"
)

func main(){

	var pivo float64
	var ml1 float64
	var d float64
	var ml2 float64
	var ml22 int 
	var ml3 float64
	var difzero [3]float64
    matriz:= [3][3]float64{}
	var y1 float64
	var y2 float64 
	var y3 float64
	var x1 float64
	var x2 float64 
	var x3 float64


	for m:=0;m<3;m++{
		fmt.Printf("digite a linha %d\n",m+1)
		for n:=0;n<3;n++{
			fmt.Scan(&matriz[m][n])
		}
	}


	for m:=0;m<3;m++{
		for n:=0;n<3; n++{

			fmt.Printf("%f ", matriz[m][n])
		}
		fmt.Printf("\n")
	}

	fmt.Println("vamos achar o pivo agora")

	for m:=0;m<3;m++{
		for n:=0;n<3;n++{
           d = matriz[m][n]

		   if d!=0{
			pivo =d
			break
		   }else{
			continue
		   }
		}
		if d!=0{
			pivo =d
			break
		   }
	}

	fmt.Printf("o valor do pivo é %f", pivo)
 
	for m:=1; m<2;m++{
		   ml1 =matriz[1][0]/pivo
	fmt.Printf("o valor do multiplicador é de %f \n", ml1)
		for n:=0; n<3;n++{
			matriz[m][n] = matriz[m][n]- (ml1 * matriz[0][n])
			fmt.Println(matriz[m][n])
		}
		
	}
	
	for m:=2; m<3;m++{
		ml2 =matriz[2][0]/pivo
 fmt.Printf("o valor do multiplicador é de %f \n", ml2)
	 for n:=0; n<3;n++{
		 matriz[m][n] = matriz[m][n]- (ml2 * matriz[0][n])
		 fmt.Println(matriz[m][n])
	 }
	 
 }


	for m:=0;m<3;m++{
		for n:=0;n<3; n++{

			fmt.Printf("%0.2f ", matriz[m][n])
		}
		fmt.Printf("\n")
	}

      for m:=1; m<3;m++{
        
		if matriz[m][0]!=0{
			difzero[m]=matriz[m][0]
		}
		if matriz[2][1]!=0{
			difzero[m]=matriz[2][1]
		}
	  }

	

for m:=0;m<3;m++{

fmt.Printf("%0.2f diferente de zero \n", difzero[m])
}

if matriz[2][0] !=0 || matriz[2][1]!=0{

	if matriz[2][0] !=0{
		ml22= 0
	}else{
		ml22= 1
	}
	
	for m:=1;m<2;m++{
		for n:=0;n<3;n++{
           d = matriz[m][n]

		   if d!=0{
			pivo =d
			break
		   }else{
			continue
		   }
		}
		if d!=0{
			pivo =d
			break
		   }
	}

	fmt.Printf("o valor do novo pivo é %f", pivo)

	for m:=2; m<3;m++{
		ml3 =matriz[m][ml22]/pivo
 fmt.Printf("o valor do multiplicador é de %f \n", ml3)
	 for n:=0; n<3;n++{
		 matriz[m][n] = matriz[m][n]- (ml3 * matriz[1][n])
		 fmt.Println(matriz[m][n])
	 }
	 
 }

 for m:=0;m<3;m++{
	for n:=0;n<3; n++{

		fmt.Printf("%0.2f ", matriz[m][n])
	}
	fmt.Printf("\n")
}

}


fmt.Println("temos que os multiplicadores da matriz L são")

fmt.Printf("1 0 0 \n")
fmt.Printf("%0.0f/%0.0f 0 0 \n",matriz[0][1], matriz[0][0])
fmt.Printf("%0.0f/%0.1f %0.2f 0 \n",matriz[2][0],matriz[0][0],ml3)

fmt.Printf("bom agora vamos fazer L X Y = B \n")

fmt.Printf("me de os 3 valores de y:\n")
fmt.Scan(&y1)
fmt.Scan(&y2)
fmt.Scan(&y3)


fmt.Printf("temos que o valor de de y1 é %0.2f, y2 %0.2f e y3 %0.2f \n",y1,y2,y3)

fmt.Printf("temos que o valor dos calculos vão ficar:\n y1 =%0.2f \n %0.2fy1 +y2 =%0.2f \n %0.2fy1 +y2 +y3 = %0.2f\n", y1, ml1, y2, ml2,y3)
y1 = y1 
y2= y2 - (ml1*y1)
y3= (y1*ml2 + y2) - y3

fmt.Printf("temos que os novos valores de de y1 é %0.2f, y2 %0.2f e y3 %0.2f \n",y1,y2,y3)
fmt.Println("agora vamos resolver a segunda parte do nosso sistema que é U * X= Y \n ")
fmt.Printf("a formula vai ficar da seguinte maneira:\n %0.2fx1 + %0.2fx2 + %0.2fx3 = %0.2f\n %0.2fx2 + %0.2fx3 = %0.2f \n %0.2fx3 = %0.2f\n", matriz[0][0],matriz[0][1],matriz[0][2],y1,matriz[1][1],matriz[1][2],y2, matriz[2][2],y3)
x3= (matriz[2][2]*x3) *y3
x2= ((matriz[2][1] * x2) + (matriz[0][1]*x3))/ y2
x1= (matriz[0][0]*x1 + matriz[0][1]*x2 + matriz[0][2]*x3)/y3
fmt.Printf("os valores de x são x1:%0.2f\n x2:%0.2f\n x3:%0.2f",x1,x2,x3)
}