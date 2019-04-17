INSERT INTO id_certifications(
    id,
    account_id,
    status,
    audit_conclusions,
    name,
    identification_number,
    id_card_type,
    id_card_start_date,
    id_card_expiry,
    address,
    sex,
    id_card_front_pic,
    id_card_back_pic,
    face_pic,
    ethnic_group,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :account_id,
    :status,
    :audit_conclusions,
    :name,
    :identification_number,
    :id_card_type,
    :id_card_start_date,
    :id_card_expiry,
    :address,
    :sex,
    :id_card_front_pic,
    :id_card_back_pic,
    :face_pic,
    :ethnic_group,
    :deleted,
    :created_at,
    :updated_at
)