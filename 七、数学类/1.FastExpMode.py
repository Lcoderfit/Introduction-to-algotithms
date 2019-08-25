class Solution:
    def fast_expMod(self, a, b, m):
        """
        求a^b%m
        """
        ret = 1
        while b != 0:
            if (b & 1) == 1:
                ret = (ret * a) % m
            b >>= 1
            a = (a * a) % m
        return ret