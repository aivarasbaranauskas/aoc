let rec pow a = function
  | 0 -> 1
  | 1 -> a
  | n ->
      let b = pow a (n / 2) in
      b * b * if n mod 2 = 0 then 1 else a

let modulo x y =
  let result = x mod y in
  if result >= 0 then result else result + y
