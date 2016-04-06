
# coding: utf-8

# In[1]:

print("Hello Python!")


# In[2]:

print("Hello 역량 강화")


# In[7]:

def hanoi(n, start, end, mid):
    if n == 1:
        print(start, "->", end)
        return
    
    hanoi(n-1, start, mid, end)
    hanoi(1, start, end, mid)
    hanoi(n-1, mid, end, start)


# In[8]:

hanoi(3, 1, 2, 3)


# In[ ]:



