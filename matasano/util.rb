# typed: true
require 'sorbet-runtime'
require 'base64'

B64_TABLE = %w[
  A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
  a b c d e f g h i j k l m n o p q r s t u v w x y z
  0 1 2 3 4 5 6 7 8 9 + /
].freeze

# Ruby implementations of matasano crypto challenge
class Util
  extend T::Sig

  sig { params(num: String, ensure_size_multiple: Integer).returns(String) }
  def hex_to_b2(num, ensure_size_multiple)
    n = num.to_i(16)
    n_base_two = n.to_s(2)
    offset = ensure_size_multiple - (n_base_two.length % ensure_size_multiple)
    offset.times do
      n_base_two = "0#{n_base_two}"
    end
    n_base_two
  end

  sig { params(num: String).returns(String) }
  def hex_to_b64(num)
    num_in_base_two = hex_to_b2(num, 6)
    num_chunks = num_in_base_two.scan(/.{6}/)
    res = ''
    num_chunks.each do |chunk|
      unless chunk.is_a? Array
        i = chunk.to_i(2)
        res = "#{res}#{B64_TABLE[i]}"
      end
    end
    res
  end

  sig { params(a: String, b: String).returns(String)}
  def fixed_xor(a, b)
    num_a = a.to_i(16)
    num_b = b.to_i(16)
    res = num_a ^ num_b
    res.to_s(16)
  end

  sig { params(str: String, len: Integer).returns(String) }
  def repeat_char_by_len(str, len)
    str * len
  end

  sig { params(str: String).returns(String) }
  def hex_str_to_ascii(str)
    [str].pack('H*')
  end
end
