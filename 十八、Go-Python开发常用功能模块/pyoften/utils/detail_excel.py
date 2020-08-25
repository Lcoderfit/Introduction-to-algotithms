import os
import xlrd
import xlwt

from xlutils.copy import copy


def read_excel(path, sheet_name):
    try:
        workbook = xlrd.open_workbook(path)
        print(workbook.sheet_names())
        sheet = workbook.sheet_by_name(sheet_name)
        rows = sheet.nrows
        cols = sheet.ncols
        print("rows: %d, clos: %d" % (rows, cols))

        # 第四行第二列开始
        pos = 2
        ret = dict()
        for i in range(1 + pos, 42 + pos):
            name = sheet.cell(i, 1)
            addr = sheet.cell(i, 2)

            name = str(name)
            addr = str(addr)
            ret[name] = addr
        return ret
    except Exception as e:
        print("read error: ", e)


# 追加
def write_append_excel(path, sheet_name, write_data):
    workbook = xlrd.open_workbook(path)
    print(workbook.sheet_names())
    sheet = workbook.sheet_by_name(sheet_name)
    rows = sheet.nrows
    cols = sheet.ncols
    print("rows: %d, cols: %d" % (rows, cols))

    name_list = list()
    pos = 3
    for i in range(pos, 41 + pos):
        name = sheet.cell(i, 1)
        name_list.append(str(name))
    print(name_list)
    print(len(name_list))

    wb = copy(workbook)
    ret_sheet = wb.get_sheet(6)
    print("ret_sheet: ", ret_sheet)
    print(write_data)
    for i in range(pos, 41 + pos):
        name = name_list[i - pos]

        wd = write_data[name].lstrip("text:'").rlstrip("'")
        print(i, name, wd)
        ret_sheet.write(i, 8, wd)
    os.remove(path)
    wb.save(r"C:\Users\Lcoderfit\Desktop\荣山小学二（2）花名册 - 副本.xls")


# 新建一个excel写入
def write_excel(path, sheet_name, write_data):
    workbook = xlrd.open_workbook(path)
    print(workbook.sheet_names())
    sheet = workbook.sheet_by_name(sheet_name)
    rows = sheet.nrows
    cols = sheet.ncols
    print("rows: %d, cols: %d" % (rows, cols))

    name_list = list()
    pos = 3
    for i in range(pos, 41 + pos):
        name = sheet.cell(i, 1)

        name = str(name)
        name_list.append(name)
    print(name_list)
    print(len(name_list))

    write_workbook = xlwt.Workbook()
    write_sheet = write_workbook.add_sheet("B")
    for i in range(41):
        name = name_list[i]
        wd = write_data[name].lstrip("text:'").rstrip("'")
        write_sheet.write(i, 0, wd)
    write_workbook.save(r"C:\Users\Lcoderfit\Desktop\test.xls")


if __name__ == "__main__":
    read_path = r"C:\Users\Lcoderfit\Desktop\荣山小学二2班学生情况再摸排表.xlsx"
    read_sheet_name = r"学生摸排表"
    ret = read_excel(read_path, read_sheet_name)

    print(ret)
    print("ret's len: ", len(ret))
    print()

    write_path = r"C:\Users\Lcoderfit\Desktop\荣山小学二（2）花名册 - 副本.xls"
    write_sheet_name = r"二（2）"
    write_excel(write_path, write_sheet_name, ret)


