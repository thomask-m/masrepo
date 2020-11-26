# frozen_string_literal: true
# typed: true

module T::Types
  class TypedArray < TypedEnumerable
    # @override Base
    def name
      "T::Array[#{@type.name}]"
    end

    def underlying_class
      Array
    end

    # @override Base
    def recursively_valid?(obj)
      obj.is_a?(Array) && super
    end

    # @override Base
    def valid?(obj)
      obj.is_a?(Array)
    end

    def new(*args)
      Array.new(*T.unsafe(args))
    end

    class Untyped < TypedArray
      def initialize
        super(T.untyped)
      end

      def valid?(obj)
        obj.is_a?(Array)
      end
    end
  end
end
