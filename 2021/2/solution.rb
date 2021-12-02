i = File.read("input").lines.map { |e| e.split(" ") }

horizontal_position = 0
depth = 0

i.each { |e|
  x = e[1].to_i
  case e[0]
  when "up"
    depth -= x
  when "down"
    depth += x
  else
    horizontal_position += x
  end
}

solution_1 = horizontal_position * depth

horizontal_position = 0
depth = 0
aim = 0

i.each { |e|
  x = e[1].to_i
  case e[0]
  when "up"
    aim -= x
  when "down"
    aim += x
  else
    horizontal_position += x
    depth += aim * x
  end
}

solution_2 = horizontal_position * depth

print "Solution for part 1: "
puts solution_1

print "Solution for part 2: "
puts solution_2
