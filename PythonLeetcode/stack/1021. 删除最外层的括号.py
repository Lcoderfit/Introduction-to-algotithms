"""

方法1： 
时间复杂度：O()
空间复杂度：O()

case1:
(()())()(())
r:
()()()()

case2:
()()
r:
""
"""
import sys
from typing import List


class Solution:
    @staticmethod
    def remove_outer_parentheses(s: str) -> str:
        if not s:
            return ""
        stack = [s[0]]
        pos = 0
        ans = ""
        for i in range(1, len(s)):
            if not stack:
                ans += s[pos + 1:i - 1]
                pos = i
                stack.append(s[i])
            else:
                if stack[-1] == "(" and s[i] == ")":
                    stack.pop()
                else:
                    stack.append(s[i])
        return ans + s[pos + 1:-1]


# if __name__ == '__main__':
#     sl = Solution()
#     s_cur = input().strip()
#     print(sl.remove_outer_parentheses(s_cur))
compay_name = \
('中国电信股份有限公司顺德大良新基客户服务中心',
'北京汉普在线文化发展有限公司深圳南山分公司',
'通辽市科尔沁区通科原粮东喜养殖专业合作社',
'乌鲁木齐市天山区解放北路社区服务中心',
'一壶浊酒(衢州)生态农业发展有限公司',
'万安县蔚兰教育咨询有限公司',
'万源市青花农贸市场(普通合伙)',
'淮南市开阳实业公司',
'洛阳六合科技开发公司',
'衡阳市城北实验试剂仪器经营部',
'哈尔滨市晟禹小额贷款股份有限公司',
'苏州市商业银行股份有限公司',
'武汉现代信达办公设备经营部',
'武汉市恒安铁路车辆维修配件厂',
'余干县车名家汽车销售服务有限公司',
'上海湃睿信息科技有限公司深圳分公司',
'玉门农垦实业公司',
'武汉市新辉纺机配件厂',
'哈尔滨森鹰窗业股份有限公司',
'常州康达工业总公司(常州厚余煤矿)',
'庄浪县悉地蔬菜种植农民专业合作社',
'曹县芳华文化传媒有限公司',
'丰城市继开物业管理有限公司',
'沈阳全密封变压器股份有限公司',
'惠州亿纬锂能股份有限公司',
'上杭县才溪农村信用合作社',
'滁州针织三厂',
'宁国市霞西医药商店',
'张掖地区丰达经销公司',
'北京海淀天星电子产品经销公司',
'四川恒泰勘察岩土工程有限责任公司',
'秦皇岛市区农村信用合作联社',
'绵阳市涪江造纸厂',
'惠州亿纬锂能股份有限公司',
'富滇银行股份有限公司',
'江西长运股份有限公司',
'安庆市外联物资公司',
'西宁市彭家寨供销合作社',
'哈尔滨森鹰窗业股份有限公司',
'呼伦贝尔市金属化建有限责任公司',
'永嘉县陶瓷面砖厂',
'四川东方老人乐园',
'元谋县农业资源综合开发有限责任公司',
'酒泉市兴鑫农机实业有限责任公司',
'元谋县农业资源综合开发有限责任公司',
'元谋县农业资源综合开发有限责任公司')


if __name__ == '__main__':
    # print("(", end="")
    # for name in compay_name:
    #     print
    print(compay_name)