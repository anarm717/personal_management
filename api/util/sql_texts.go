package util

const Sql_getEmploees ="SELECT he.ID,cp.SURNAME , cp.NAME , cp.PATRONYMIC , cp.BIRTH_DATE as birthdate FROM humres_kadr.HR_EMPLOYEE he , COMMON_KADR.COM_PERSON cp WHERE he.PERSON_ID = cp.ID and rownum<50 ORDER BY he.id OFFSET 0 ROWS FETCH NEXT 14 ROWS ONLY"
