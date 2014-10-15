  SRC_ROOT		= File.join(File.expand_path("."), "src", "exercises-in-programming-style")
  PRIDE_AND_PREJUDICE	= File.join(SRC_ROOT, "pride-and-prejudice.txt")
  STOP_WORDS		= File.join(SRC_ROOT, 'stop_words.txt')
  @stop_words = File.read(STOP_WORDS).split(',')
  @stop_words << ('a'..'z').to_a # also alphabet
  @stop_words.flatten!.uniq!

  @words = Hash.new {|h,k| h[k] = 0 }


  File.read(PRIDE_AND_PREJUDICE).each_line do |line|
    line.gsub!(/[^a-zA-Z0-9]/, " ") # remove non alphanumeric
    words = line.split(" ")
    words.each do |w|
      next if @stop_words.include?(w.downcase)
      @words[w.downcase] += 1
    end
  end

  @words.sort {|a,b| a[1] <=> b[1]}.reverse[0...25].each do |k, v|
    puts "#{k}  -  #{v}"
  end
