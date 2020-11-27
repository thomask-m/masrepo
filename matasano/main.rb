require_relative 'util'

code = '1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736'
util = Util.new
('A'..'Z').each do |alpha|
  alpha.each_byte do |c|
    res = util.fixed_xor(util.repeat_char_by_len(c.to_s(16), code.length/2), code)
    p util.hex_str_to_ascii(res)
  end
end

# message is "Cooking MC's like a pound of bacon"
