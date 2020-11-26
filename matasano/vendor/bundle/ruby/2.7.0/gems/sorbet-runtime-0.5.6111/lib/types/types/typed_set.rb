# frozen_string_literal: true
# typed: true

module T::Types
  class TypedSet < TypedEnumerable
    attr_reader :type

    def underlying_class
      Hash
    end

    # @override Base
    def name
      "T::Set[#{@type.name}]"
    end

    # @override Base
    def recursively_valid?(obj)
      obj.is_a?(Set) && super
    end

    # @override Base
    def valid?(obj)
      obj.is_a?(Set)
    end

    def new(*args)
      Set.new(*T.unsafe(args))
    end

    class Untyped < TypedSet
      def initialize
        super(T.untyped)
      end

      def valid?(obj)
        obj.is_a?(Set)
      end
    end
  end
end
