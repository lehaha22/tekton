# tekton k8s

## 统计字符串出现次数最多的字符,并统计其次数
- 返回是map报**panic: assignment to entry in nil map**

  ```go
  result[key] = map1[key]
  在声明result后并未初始化它，所以它的值是nil, 不指向任何内存地址。需要通过make方法分配确定的内存地;
  result = make(map[int]int)
  注意：同为引用类型的slice，在使用append 向nil slice追加新元素就可以，原因是append方法在底层为slice重新分配了相关数组让nil slice指向了具体的内存地址
  ```

- test时传值要使用单引号(rune类型)

