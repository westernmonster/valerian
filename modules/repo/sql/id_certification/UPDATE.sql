UPDATE id_certifications
SET
    account_id=:account_id,
    status=:status,
    audit_conclusions=:audit_conclusions,
    name=:name,
    identification_number=:identification_number,
    id_card_type=:id_card_type,
    id_card_start_date=:id_card_start_date,
    id_card_expiry=:id_card_expiry,
    address=:address,
    sex=:sex,
    id_card_front_pic=:id_card_front_pic,
    id_card_back_pic=:id_card_back_pic,
    face_pic=:face_pic,
    ethnic_group=:ethnic_group,
    updated_at=:updated_at
WHERE id=:id
