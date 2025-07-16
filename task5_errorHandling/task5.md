### 🔹 **Задача 5. Обработка ошибок (время: \~1 ч)**

**Условие:**
Реализуйте функцию `Divide(a, b float64) (float64, error)`, которая возвращает ошибку при делении на 0. Используйте `errors.New`.

**Тест:**

```
fmt.Println(Divide(10, 2)) // 5.0, nil
fmt.Println(Divide(10, 0)) // 0.0, error: division by zero
```
