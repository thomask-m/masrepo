# typed: strict
require_relative 'util'
util = Util.new
max = 0
result = ''
File.foreach('codes.txt') do |code|
  ('A'..'Z').each do |alpha|
    alpha.each_byte do |c|
      res = util.fixed_xor(util.repeat_char_by_len(c.to_s(16), code.length/2), code)
      ascii_res = util.hex_str_to_ascii(res)
      score = util.evaluate_code(res)
      if score > max
        max = score
        result = ascii_res
      end
    end
  end
end
p result
p max
# code = '1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736'
# set1 challenge3: message is "Cooking MC's like a pound of bacon"
