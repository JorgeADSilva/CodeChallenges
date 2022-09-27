#ifndef EMPLOYEE_FUNCTIONS
#define EMPLOYEE_FUNCTIONS

typedef struct
{
    int day, month, year;
} DATA;

typedef struct
{
    int ID;
    char name[30];
    DATA birthday;
    float salary;
} EMPLOYEE;

// helper
void printEmployeesIds(EMPLOYEE *employeeArray, int numberOfEmployees)
{
    printf("[");
    for (int i = 0; i < numberOfEmployees; i++)
    {
        printf("%i", i);
        if (i + 1 != numberOfEmployees)
            printf(", ");
    }
    printf("]");
}

// OPTION 1
void insertNewEmployee(int *numberOfEmployees, EMPLOYEE *employeeArray)
{
    EMPLOYEE employeeForm;
    printf("---EMPLOYEE-FORM---\n\n");
    printf("Name: ");
    scanf("%s", &employeeForm.name);

    printf("Salary: ");
    scanf("%f", &employeeForm.salary);

    printf("Birthday Date (DD/MM/YYYY FORMAT): ");
    scanf("%d/%d/%d/", &employeeForm.birthday.day, &employeeForm.birthday.month, &employeeForm.birthday.year);
    employeeArray[*numberOfEmployees] = employeeForm;
    *numberOfEmployees += 1;
    printf("-------------------\n\n");
}
// OPTION 2
void doubleEmployeeSalaryById(EMPLOYEE *employeeArray, int numberOfEmployees)
{
    if (numberOfEmployees == 0)
    {
        printf("Not Possible for now!\n");
        return;
    }
    int id;
    printf("Available IDs: ");
    printEmployeesIds(employeeArray, numberOfEmployees);
    printf("\nWhat is the employee ID? ");
    scanf("%d", &id);
    if (id > numberOfEmployees - 1)
    {
        printf("ID out of range, operation canceled.\n");
    }
    else
        employeeArray[id].salary = 2 * employeeArray[id].salary;
}
// OPTION 3
void printListOfEmployees(float salary, EMPLOYEE *employeeArray, int numberOfEmployees)
{
    if (numberOfEmployees == 0)
    {
        printf("No inserted data yet\n");
        return;
    }
    printf("List of employees with salary bellow to %.2f (euros):\n", salary);
    printf("|\tID\t|\tName\t|\tBirthday\t|\tSalary (euros)\t|\n");
    printf("|\t-----\t|\t----\t|\t--------\t|\t--------------\t|\n");
    for (int i = 0; i < numberOfEmployees; i++)
    {
        if (employeeArray[i].salary < salary)
        {
            printf("|\t%i\t|\t%s\t|\t %d/%d/%d \t|\t %.2f euros\t|\n", i, employeeArray[i].name, employeeArray[i].birthday.day, employeeArray[i].birthday.month, employeeArray[i].birthday.year, employeeArray[i].salary);
            printf("|\t-----\t|\t----\t|\t--------\t|\t--------------\t|\n");
        }
    }
}
#endif
