select w2.id
from Weather w1, Weather w2
where w1.Temperature  < w2.Temperature and (TO_DAYS(w2.recordDate)-TO_DAYS(w1.recordDate))=1

