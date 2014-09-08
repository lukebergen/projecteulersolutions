("001".."477").each do |num|
  `curl https://projecteuler.net/problem=#{num} > #{num}/problem.html`
  str = File.read("#{num}/problem.html")
  str = "<h2>" + str.split('<h2>')[1..-1].join('')
  str = str.split('<div id="footer" class="noprint">')[0]
  str = "<html>
#{str}
</html>"
  File.open("#{num}/problem.html", 'w') {|f| f.write(str)}
end
