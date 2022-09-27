#include <stdio.h>
#include "employee_functions.h"
#include <locale.h>
#define MAX 1000

int main()
{
    EMPLOYEE employeesArray[MAX];
    int N = 0;
    int menuOption = 0;
    while (menuOption != 4)
    {

        printf("\t----------------------------------------MENU----------------------------------------\n");
        printf("\t (1) - Insert new employee\n");
        printf("\t (2) - Double employee salary by ID\n");
        printf("\t (3) - List employees with salaries bellow to 1000 euros\n");
        printf("\t (4) - Sair\n");
        printf("\t------------------------------------------------------------------------------------\n");
        printf("\t Pick option: ");
        scanf("%d", &menuOption);
        printf("\n\n");

        switch (menuOption)
        {
        case 1:
            insertNewEmployee(&N, employeesArray);
            break;
        case 2:
            doubleEmployeeSalaryById(employeesArray, N);
            break;
        case 3:
            printListOfEmployees(1000, employeesArray, N);
            break;
        case 4:
            break;
        }
    }
}