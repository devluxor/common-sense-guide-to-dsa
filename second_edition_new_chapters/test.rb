def count_x(string, count=0)
  return 0 if string.length === 0 
  
  if string[0] === 'x'
    return 1 + count_x(string[1..])
  end
  
  count_x(string[1..])
end

puts count_x("axbxcxd")