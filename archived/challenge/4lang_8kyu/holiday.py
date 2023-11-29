def duty_free(price, discount, holiday_cost):
    print(int(holiday_cost//(price*discount/100))) 

duty_free(12, 50, 1000)# 166
duty_free(17, 10, 500)# 294

