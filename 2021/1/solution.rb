i = File.read("input").lines.map(&:to_i)

print "Solution for part 1: "
puts i.each_cons(2).count { |a,b| b > a }

print "Solution for part 2: "
puts i.each_cons(3).map { |e| e.sum }.each_cons(2).count { |a,b| b > a }
